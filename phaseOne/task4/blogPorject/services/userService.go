package services

import (
	"errors"
	"gotask/phaseOne/task4/blogPorject/config"
	"gotask/phaseOne/task4/blogPorject/structs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key")

func Login(username, password string) (string, error) {
	var user structs.User
	if err := config.DB.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return "", errors.New("用户名或密码错误")
	}
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errors.New("生成token失败")
	}
	return tokenStr, nil
}
