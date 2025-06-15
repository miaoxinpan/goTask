package handlers

import (
	"gotask/phaseOne/task4/blogPorject/services"
	"gotask/phaseOne/task4/blogPorject/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", Login)       // 登录
		userGroup.POST("/register", Register) // 注册
	}
}

// 登录
func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		structs.RespondWithResult(c, http.StatusBadRequest, "参数错误", nil)
		return
	}
	token, err := services.Login(req.Username, req.Password)
	if err != nil {
		structs.RespondWithResult(c, http.StatusUnauthorized, err.Error(), nil)
		return
	}
	structs.RespondWithResult(c, http.StatusOK, "登录成功！", gin.H{"token": token})
}

// 注册
func Register(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		structs.RespondWithResult(c, http.StatusBadRequest, "参数错误", nil)
		return
	}
	err := services.Register(req.Username, req.Password, req.Email)
	if err != nil {
		structs.RespondWithResult(c, http.StatusUnauthorized, err.Error(), nil)
		return
	}
	structs.RespondWithResult(c, http.StatusOK, "注册成功！", nil)

}
