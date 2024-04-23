package dto

import "gin-app/models"

type UserInfoDto struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Email      string `json:"email" `
	Avatar     string `json:"avatar"`
	LoginTime  string `json:"login_time"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func ToUserInfoDto(user models.User) UserInfoDto {
	return UserInfoDto{
		ID:         user.ID,
		Name:       user.Name,
		Phone:      user.Phone,
		Email:      user.Email,
		Avatar:     user.Avatar,
		LoginTime:  user.LoginTime,
		CreateTime: user.CreateTime,
		UpdateTime: user.UpdateTime,
	}
}
