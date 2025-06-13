package router

import (
	"gotask/phaseOne/task4/blogPorject/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	handlers.RegisterUserRoutes(r)
	handlers.RegisterPostRoutes(r)
	handlers.RegisterCommentRoutes(r)
	return r
}
