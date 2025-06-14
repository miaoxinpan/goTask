package middleware

import (
	"gotask/phaseOne/task4/blogPorject/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

var jwtSecret = []byte(viper.GetString("jwt.secret"))

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			structs.RespondWithResult(c, http.StatusUnauthorized, "未提供Token", nil)
			c.Abort()
			return
		}
		// Authorization: Bearer <你的token>  直接忽略这个Bearer
		// parts := strings.SplitN(authHeader, " ", 2)
		// if !(len(parts) == 2 && parts[0] == "Bearer") {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Token格式错误"})
		// 	c.Abort()
		// 	return
		// }
		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			structs.RespondWithResult(c, http.StatusUnauthorized, "无效的Token", nil)
			c.Abort()
			return
		}
		// 解析 claims 并写入 gin.Context
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if userID, ok := claims["user_id"]; ok {
				c.Set("user_id", userID)
			}
		}
		c.Next()
	}
}

func GetJWTKey() []byte {
	return jwtSecret
}
