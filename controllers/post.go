package controllers

import (
	"gin-app/common/code"
	"gin-app/common/global"
	"gin-app/common/response"
	"gin-app/dto"
	"gin-app/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {
	json := dto.CreatePostReq{}

	if err := ctx.ShouldBindJSON(&json); err != nil {
		response.ReturnError(ctx, code.CommonError, err.Error())
		return
	}

	user, _ := ctx.Get("user")

	post := models.Post{
		UserId:     user.(models.User).ID,
		CategoryId: json.CategoryId,
		Title:      json.Title,
		Content:    json.Content,
		Image:      json.Image,
		CreateTime: time.Now().Format(time.DateTime),
		UpdateTime: time.Now().Format(time.DateTime),
	}

	global.DB.Create(&post)

	response.ReturnOK(ctx, post)
}

func GetPostInfo(ctx *gin.Context) {
	postId := ctx.Query("postId")

	var post models.Post
	global.DB.Preload("Category").Where("id = ?", postId).First(&post)

	if post.ID == 0 {
		response.ReturnError(ctx, code.CommonError, "文章不存在")
		return
	}

	response.ReturnOK(ctx, post)
}

func GetPostList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "20"))

	page = page - 1

	var posts []dto.GetPostListRes
	global.DB.Table("post").
		Order("create_time desc").
		Offset(page * size).
		Limit(size).
		Select("post.id, post.user_id, post.category_id, category.name as category_name, post.title, post.content, post.image, post.update_time, post.create_time").
		Joins("left join category on post.category_id = category.id").
		Scan(&posts)

	var total int64
	global.DB.Model(models.Post{}).Count(&total)

	response.ReturnList(ctx, posts, total)
}
