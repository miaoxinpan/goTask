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
		postGroup.POST("/createpost", middleware.JWTAuthMiddleware(), CreatePost)                   // 发表文章
		postGroup.GET("/getpostforid", middleware.JWTAuthMiddleware(), GetPostForId)                // 获取单个文章的详细信息(根据文章id来)
		postGroup.POST("/updatepostforauthor", middleware.JWTAuthMiddleware(), UpdatePostForAuthor) //实现文章的更新功能，实现文章的删除功能。只有文章的作者才能操作自己的文章。
		postGroup.GET("/getpostforuserid", middleware.JWTAuthMiddleware(), GetPostForUserId)        // 根据用户id获取用户发表过的文章
		postGroup.GET("/all", GetAllPost)                                                           // 获取全部文章 根据发布时间排序 需要做分页
	}
}

// 发表文章
func CreatePost(c *gin.Context) {
	var post structs.Post = structs.Post{}
	if err := c.ShouldBindJSON(&post); err != nil {
		structs.RespondWithResult(c, http.StatusBadRequest, "参数错误", nil)
		return
	}
	userID, exists := c.Get("user_id")
	if !exists {
		structs.RespondWithResult(c, http.StatusUnauthorized, "用户信息获取失败", nil)
		return
	}
	//post.UserID = uint(userID)  直接转报错
	//通常通过 JWT 解析出来的 user_id 是 float64 类型，所以你需要先断言为 float64，再转换为 uint。
	intUserId, ok := userID.(float64)
	if !ok {
		structs.RespondWithResult(c, http.StatusUnauthorized, "用户ID类型错误", nil)
		return
	}

	//把userId放到post里面去
	post.UserID = uint(intUserId)
	err := services.CreatePost(post)
	if err != nil {
		structs.RespondWithResult(c, http.StatusUnauthorized, err.Error(), nil)
		return
	}
	structs.RespondWithResult(c, http.StatusOK, "发布成功！", nil)
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
		structs.RespondWithResult(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	structs.RespondWithResult(c, http.StatusOK, "查询成功", gin.H{
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
		structs.RespondWithResult(c, http.StatusUnauthorized, "参数错误,请检查参数!", nil)
		return
	}
	_, err := fmt.Sscanf(p, "%d", &postId)
	if err != nil {
		structs.RespondWithResult(c, http.StatusBadRequest, "postId参数格式错误!", nil)
		return
	}
	var post structs.Post
	err = services.GetPostForId(postId, &post)
	if err != nil {
		structs.RespondWithResult(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	structs.RespondWithResult(c, http.StatusOK, "查询成功！", post)
}

//实现文章的更新功能，实现文章的删除功能。只有文章的作者才能操作自己的文章。
//根据postid 只能更新文章内容 校验userid是不是登录的的这个id

func UpdatePostForAuthor(c *gin.Context) {
	var req struct {
		Postid  uint   `json:"postid" binding:"required"`
		Content string `json:"content" `
		UserID  uint   `json:"userid" binding:"required"`
		OpType  string `json:"opType" binding:"required"` //D删除  U 更新
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		structs.RespondWithResult(c, http.StatusBadRequest, "参数错误", nil)
		return
	}
	if req.OpType == "U" && req.Content == "" {
		structs.RespondWithResult(c, 400, "更新操作时 content 不能为空", nil)
		return
	}
	userID, exists := c.Get("user_id")
	if !exists {
		structs.RespondWithResult(c, http.StatusUnauthorized, "用户信息获取失败", nil)
		return
	}
	//post.UserID = uint(userID)  直接转报错
	//通常通过 JWT 解析出来的 user_id 是 float64 类型，所以你需要先断言为 float64，再转换为 uint。
	currentUserId, ok := userID.(float64)
	if !ok {
		structs.RespondWithResult(c, http.StatusUnauthorized, "用户ID类型错误", nil)
		return
	}
	if req.UserID != uint(currentUserId) {
		structs.RespondWithResult(c, http.StatusUnauthorized, "您没有权限更新此文章!", nil)
		return
	}
	err := services.UpdatePostForAuthor(req.Postid, req.Content, req.OpType)
	if err != nil {
		structs.RespondWithResult(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	structs.RespondWithResult(c, http.StatusOK, "更新成功！", nil)
}

// 根据用户id获取用户发表过的文章
func GetPostForUserId(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		structs.RespondWithResult(c, http.StatusUnauthorized, "用户信息获取失败", nil)
		return
	}
	//post.UserID = uint(userID)  直接转报错
	//通常通过 JWT 解析出来的 user_id 是 float64 类型，所以你需要先断言为 float64，再转换为 uint。
	id, ok := userID.(float64)
	if !ok {
		structs.RespondWithResult(c, http.StatusUnauthorized, "用户ID类型错误", nil)
		return
	}
	var posts []structs.Post
	err := services.GetPostForUserId(uint(id), &posts)
	if err != nil {
		structs.RespondWithResult(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	structs.RespondWithResult(c, http.StatusOK, "查询成功！", posts)

}
