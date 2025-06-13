package handlers

import "github.com/gin-gonic/gin"

func RegisterCommentRoutes(r *gin.Engine) {
	commentGroup := r.Group("/comment")
	{
		commentGroup.POST("/", CreateComment)
		commentGroup.GET("/:id", GetComment)
		// ...更多文章相关接口
	}
}

// 发表评论
func CreateComment(c *gin.Context) { /* ... */ }

//获取评论
func GetComment(c *gin.Context) { /* ... */ }
