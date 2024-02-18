package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
	"github.com/wneessen/go-mail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type UserService struct {
	container framework.Container
	logger    contract.Log
	configer  contract.Config
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func genToken(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (u *UserService) Register(ctx context.Context, user *User) (*User, error) {
	ormService := u.container.MustMake(contract.ORMKey).(contract.ORMService)
	db, err := ormService.GetDB()
	if err != nil {
		return nil, err
	}
	userDB := &User{}
	if db.Where(&User{Email: user.Email}).First(userDB).Error != gorm.ErrRecordNotFound {
		return nil, errors.New("邮箱已注册用户，不能重复注册")
	}
	if db.Where(&User{UserName: user.UserName}).First(userDB).Error != gorm.ErrRecordNotFound {
		return nil, errors.New("用户名已经被注册，请换一个用户名")
	}
	token := genToken(10)
	user.Token = token

	// 将请求注册进入缓存，保存一天
	cacheService := u.container.MustMake(contract.CacheKey).(contract.CacheService)
	key := fmt.Sprintf("user:register:%v", user.Token)
	if err := cacheService.SetObj(ctx, key, user, 24*time.Hour); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) SendRegisterMail(ctx context.Context, user *User) error {
	logger := u.container.MustMake(contract.LogKey).(contract.Log)
	configer := u.container.MustMake(contract.ConfigKey).(contract.Config)

	// 配置服务中获取发送邮件需要的参数
	host := configer.GetString("app.smtp.host")
	port := configer.GetInt("app.smtp.port")
	username := configer.GetString("app.smtp.username")
	password := configer.GetString("app.smtp.password")
	from := configer.GetString("app.smtp.from")
	domain := configer.GetString("app.domain")

	// 组装message
	m := mail.NewMsg()
	m.Subject("感谢您注册我们的bbs")

	link := fmt.Sprintf("%v/api/user/register/verify?token=%v", domain, user.Token)
	m.SetBodyString(mail.TypeTextHTML, fmt.Sprintf("请点击下面的链接完成注册：%s", link))

	if err := m.From(from); err != nil {
		logger.Error(ctx, "set from error", map[string]interface{}{
			"err":     err,
			"message": m,
		})
		return err
	}
	if err := m.To(user.Email); err != nil {
		logger.Error(ctx, "set to error", map[string]interface{}{
			"err":     err,
			"message": m},
		)
		return err
	}

	// 发送电子邮件
	c, err := mail.NewClient(host, mail.WithPort(port), mail.WithSMTPAuth(mail.SMTPAuthLogin), mail.WithUsername(username), mail.WithPassword(password))
	if err != nil {
		logger.Error(ctx, "create mail client error", map[string]interface{}{
			"err": err,
		})
		return err
	}
	if err := c.DialAndSend(m); err != nil {
		logger.Error(ctx, "send email error", map[string]interface{}{
			"err":     err,
			"message": m,
		})
		return err
	}
	return nil
}

func (u *UserService) VerifyRegister(ctx context.Context, token string) (bool, error) {
	//验证token
	cacheService := u.container.MustMake(contract.CacheKey).(contract.CacheService)
	key := fmt.Sprintf("user:register:%v", token)
	user := &User{}
	if err := cacheService.GetObj(ctx, key, user); err != nil {
		return false, err
	}
	if user.Token != token {
		return false, nil
	}

	//验证邮箱，用户名的唯一
	ormService := u.container.MustMake(contract.ORMKey).(contract.ORMService)
	db, err := ormService.GetDB()
	if err != nil {
		return false, err
	}
	userDB := &User{}
	if !errors.Is(db.Where(&User{Email: user.Email}).First(userDB).Error, gorm.ErrRecordNotFound) {
		return false, errors.New("邮箱已注册用户，不能重复注册")
	}
	if !errors.Is(db.Where(&User{UserName: user.UserName}).First(userDB).Error, gorm.ErrRecordNotFound) {
		return false, errors.New("用户名已经被注册，请换一个用户名")
	}

	// 密码加密存储
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return false, err
	}
	user.Password = string(hash)
	// 创建用户
	if err := db.Create(user).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserService) Login(ctx context.Context, user *User) (*User, error) {
	ormService := u.container.MustMake(contract.ORMKey).(contract.ORMService)
	db, err := ormService.GetDB()
	if err != nil {
		return nil, err
	}
	userDB := &User{}
	if err := db.Where("username=?", user.UserName).First(userDB).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password)); err != nil {
		return nil, err
	}

	userDB.Password = ""
	//缓存session
	cacheService := u.container.MustMake(contract.CacheKey).(contract.CacheService)
	token := genToken(10)
	key := fmt.Sprintf("session:%v", token)
	if err := cacheService.SetObj(ctx, key, userDB, 24*time.Hour); err != nil {
		return nil, err
	}
	userDB.Token = token

	return userDB, nil
}

func (u *UserService) Logout(ctx context.Context, user *User) error {
	cacheService := u.container.MustMake(contract.CacheKey).(contract.CacheService)
	userSession, err := u.VerifyLogin(ctx, user.Token)
	//不需要任何操作
	if err != nil || userSession.UserName != user.UserName {
		return nil
	}

	key := fmt.Sprintf("session:%v", user.Token)
	_ = cacheService.Del(ctx, key)
	return nil
}

func (u *UserService) VerifyLogin(ctx context.Context, token string) (*User, error) {
	if token == "" {
		return nil, errors.New("token不能为空")
	}

	cacheService := u.container.MustMake(contract.CacheKey).(contract.CacheService)
	key := fmt.Sprintf("session:%v", token)
	user := &User{}
	err := cacheService.GetObj(ctx, key, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) GetUser(ctx context.Context, userID int64) (*User, error) {
	ormService := u.container.MustMake(contract.ORMKey).(contract.ORMService)
	db, err := ormService.GetDB()
	if err != nil {
		return nil, err
	}
	user := &User{ID: userID}
	if err := db.WithContext(ctx).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserService(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	logger := container.MustMake(contract.LogKey).(contract.Log)
	configer := container.MustMake(contract.ConfigKey).(contract.Config)
	return &UserService{container: container, logger: logger, configer: configer}, nil
}
