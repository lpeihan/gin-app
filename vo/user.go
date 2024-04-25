package vo

type RegisterRequestJson struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone"`
	Email    string `json:"email" `
	Avatar   string `json:"avatar"`
}
