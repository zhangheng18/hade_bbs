package user

import (
	"context"
	"encoding/json"
	"time"
)

const UserKey = "user"

type Service interface {
	// 请在这里定义你的方法
	// Register 注册用户,注意这里只是将用户注册, 并没有激活, 需要调用
	// 参数：user必填，username，password, email
	// 返回值： user 带上token
	Register(ctx context.Context, user *User) (*User, error)
	// SendRegisterMail 发送注册的邮件
	// 参数：user必填： username, password, email, token
	SendRegisterMail(ctx context.Context, user *User) error
	// VerifyRegister 注册用户，验证注册信息, 返回验证是否成功
	VerifyRegister(ctx context.Context, token string) (bool, error)

	// Login 登录相关，使用用户名密码登录，获取完成User信息
	Login(ctx context.Context, user *User) (*User, error)
	// Logout 登出
	Logout(ctx context.Context, user *User) error
	// VerifyLogin 登录验证
	VerifyLogin(ctx context.Context, token string) (*User, error)

	// GetUser 获取用户信息
	GetUser(ctx context.Context, userID int64) (*User, error)
}

type User struct {
	ID        int64     `gorm:"column:id;primary_key"`
	UserName  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Email     string    `gorm:"column:email"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;"`

	Token string `gorm:"-"` // token 可以用作注册token或者登录token
}

// MarshalBinary 实现BinaryMarshaler 接口
func (b *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(b)
}

// UnmarshalBinary 实现 BinaryUnMarshaler 接口
func (b *User) UnmarshalBinary(bt []byte) error {
	return json.Unmarshal(bt, b)
}
