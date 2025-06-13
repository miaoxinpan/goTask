package handlers

import (
	"fmt"
	middleware "gotask/phaseOne/task4/blogPorject/middlewar"
	"gotask/phaseOne/task4/blogPorject/services"
	"gotask/phaseOne/task4/blogPorject/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(r *gin.Engine) {
	postGroup := r.Group("/post")
	{
		postGroup.POST("/createpost", middleware.JWTAuthMiddleware(), CreatePost)
		postGroup.GET("/getpostforId", middleware.JWTAuthMiddleware(), GetPostForId)
		postGroup.GET("/all", GetAllPost)
	}
}

// 发表文章
func CreatePost(c *gin.Context) {
	var post structs.Post = structs.Post{}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户信息获取失败"})
		return
	}
	//post.UserID = uint(userID)  直接转报错
	//通常通过 JWT 解析出来的 user_id 是 float64 类型，所以你需要先断言为 float64，再转换为 uint。
	intUserId, ok := userID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户ID类型错误"})
		return
	}

	//把userId放到post里面去
	post.UserID = uint(intUserId)
	err := services.CreatePost(post)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "发布成功",
	})
}

// 获取全部文章 根据发布时间排序 需要做分页
func GetAllPost(c *gin.Context) {
	//默认1 10  第一个  每页10条数据  如果传了2 5 就是要看第二页  每页5条
	page := 1
	pageSize := 10
	if p := c.Query("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}
	if ps := c.Query("pageSize"); ps != "" {
		fmt.Sscanf(ps, "%d", &pageSize)
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	var posts []structs.Post
	var total int64
	err := services.GetAllPost(page, pageSize, &posts, &total)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":     posts,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

// 获取单个文章的详细信息(根据文章id来)
func GetPostForId(c *gin.Context) {
	p := c.Query("postId")
	var postId uint
	if p == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "参数错误,请检查参数!"})
		return
	}
	_, err := fmt.Sscanf(p, "%d", &postId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "postId参数格式错误!"})
		return
	}
	var post structs.Post
	err = services.GetPostForId(postId, &post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "查询成功！",
		"data":    post,
	})
}
