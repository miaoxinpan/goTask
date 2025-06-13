package handlers

import (
	"gotask/phaseOne/task4/blogPorject/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(r *gin.Engine) {
	postGroup := r.Group("/post")
	{
		postGroup.POST("/createPost", CreatePost)
		postGroup.GET("/:id", GetPost)
	}
}

// 发表文章
func CreatePost(c *gin.Context) {
	var post structs.Post = structs.Post{}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
}

// 获取文章
func GetPost(c *gin.Context) { /* ... */ }
