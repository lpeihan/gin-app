package models

import (
	"gin-app/common/global"

	"gorm.io/gorm"
)

type Category struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Name       string `json:"name"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

func (table *Category) TableName() string {
	return "category"
}

func CreateCategory(category *Category) *gorm.DB {
	return global.DB.Create(category)
}

func GetCategoryByName(name string) Category {
	category := Category{}

	global.DB.Where("name = ?", name).First(&category)

	return category
}
