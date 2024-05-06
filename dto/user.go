package dto

type RegisterReq struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone"`
	Email    string `json:"email" `
	Avatar   string `json:"avatar"`
}

type UserInfoRes struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Email      string `json:"email" `
	Avatar     string `json:"avatar"`
	LoginTime  string `json:"loginTime"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	PostCount  int64  `json:"postCount"`
}
