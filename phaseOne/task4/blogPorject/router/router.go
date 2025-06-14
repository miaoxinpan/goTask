package router

import (
	"gotask/phaseOne/task4/blogPorject/handlers"
	middleware "gotask/phaseOne/task4/blogPorject/middlewar"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.LogRequestMiddleware()) // 日志中间件应放在最前面
	handlers.RegisterUserRoutes(r)
	handlers.RegisterPostRoutes(r)
	handlers.RegisterCommentRoutes(r)
	return r
}
