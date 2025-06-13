package handlers

import (
	"gotask/phaseOne/task4/blogPorject/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", Login)
		userGroup.POST("/register", Register)
	}
}

// 登录
func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	token, err := services.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功！",
		"token":   token,
	})
}

// 注册
func Register(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	err := services.Register(req.Username, req.Password, req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功！",
	})
}
