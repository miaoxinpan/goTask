package handlers

import (
	"fmt"
	middleware "gotask/phaseOne/task4/blogPorject/middlewar"
	"gotask/phaseOne/task4/blogPorject/services"
	"gotask/phaseOne/task4/blogPorject/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterCommentRoutes(r *gin.Engine) {
	commentGroup := r.Group("/comment")
	{
		commentGroup.POST("/createcomment", middleware.JWTAuthMiddleware(), CreateComment)            // 发表评论
		commentGroup.GET("/getcommentforuserid", middleware.JWTAuthMiddleware(), GetCommentForUserId) // 查看自己发表的评论
		commentGroup.GET("/getcommentforpostid", GetCommentForPostId)                                 // 获取某个文章全部评论

	}
}

// 发表评论
func CreateComment(c *gin.Context) {
	var req struct {
		PostID  uint   `json:"postid" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
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
	id, ok := userID.(float64)
	if !ok {
		structs.RespondWithResult(c, http.StatusUnauthorized, "用户ID类型错误", nil)
		return
	}
	err := services.CreateComment(req.PostID, req.Content, uint(id))
	if err != nil {
		structs.RespondWithResult(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	structs.RespondWithResult(c, http.StatusOK, "评论成功！", nil)
}

// 查看自己发表的评论
func GetCommentForUserId(c *gin.Context) {
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
	var comments []structs.Comment
	err := services.GetCommentForUserId(uint(id), &comments)
	if err != nil {
		structs.RespondWithResult(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	structs.RespondWithResult(c, http.StatusOK, "评论成功！", comments)

}

// 获取某个文章全部评论

func GetCommentForPostId(c *gin.Context) {
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
	var comments []structs.Comment
	err = services.GetCommentForPostId(postId, &comments)
	if err != nil {
		structs.RespondWithResult(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	structs.RespondWithResult(c, http.StatusOK, "查询成功！", comments)
}
