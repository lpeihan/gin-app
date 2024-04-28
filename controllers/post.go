package controllers

import (
	"gin-app/common/code"
	"gin-app/common/global"
	"gin-app/common/response"
	"gin-app/models"
	"gin-app/vo"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {
	json := vo.CreatePostRequestJson{}

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

	var posts []models.Post
	global.DB.Order("create_time desc").Offset(page * size).Limit(size).Find(&posts)
	// var posts = global.DB.Exec("select * from post order by create_time desc limit ? offset ?", pageSize, pageNum*pageSize)

	var total int64
	global.DB.Model(models.Post{}).Count(&total)

	response.ReturnList(ctx, posts, total)
}
