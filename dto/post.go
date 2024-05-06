package dto

type CreatePostReq struct {
	CategoryId uint   `json:"categoryId" binding:"required"`
	Title      string `json:"title" binding:"required,max=10"`
	Image      string `json:"image"`
	Content    string `json:"content" binding:"required"`
}

type GetPostListRes struct {
	ID           uint   `json:"id"`
	UserId       uint   `json:"userId"`
	CategoryId   uint   `json:"categoryId"`
	CategoryName string `json:"categoryName"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Image        string `json:"image"`
	CreateTime   string `json:"createTime"`
	UpdateTime   string `json:"updateTime"`
}
