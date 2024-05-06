package dto

type CreatePostReq struct {
	CategoryId uint   `json:"categoryId" binding:"required"`
	Title      string `json:"title" binding:"required,max=10"`
	Image      string `json:"image"`
	Content    string `json:"content" binding:"required"`
}
