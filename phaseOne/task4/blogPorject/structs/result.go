package structs

import "github.com/gin-gonic/gin"

//定义一个统一的返回格式
type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func RespondWithResult(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, Result{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
