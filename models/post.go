package models

type Post struct {
	ID         uint `json:"id" gorm:"primary_key"`
	UserId     uint `json:"userId" gorm:"not null"`
	CategoryId uint `json:"categoryId"`
	Category   *Category
	Title      string `json:"title" gorm:"type:varchar(50);not null"`
	Content    string `json:"content" gorm:"type:text;not null"`
	Image      string `json:"image"`
	Name       string `json:"name"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

func (table *Post) TableName() string {
	return "post"
}
