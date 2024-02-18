package user

import "bbs/app/provider/user"

func ConvertUserToDTO(user *user.User) *UserDTO {
	if user == nil {
		return nil
	}
	return &UserDTO{
		ID:        user.ID,
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt,
	}
}
