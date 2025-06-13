package services

import (
	"errors"
	"gotask/phaseOne/task4/blogPorject/config"
	middleware "gotask/phaseOne/task4/blogPorject/middlewar"
	"gotask/phaseOne/task4/blogPorject/structs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Login(username, password string) (string, error) {
	var user structs.User
	if err := config.DB.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return "", errors.New("用户名或密码错误")
	}
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(30 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(middleware.GetJWTKey())
	if err != nil {
		return "", errors.New("生成token失败")
	}
	return tokenStr, nil
}

func Register(username, password, email string) error {
	//先用一个sql 查询 是否有同名的或者同一个email的
	var users []structs.User
	err := config.DB.Where(" username = ? or email = ?", username, email).Find(&users)
	if err.Error != nil {
		return err.Error
	}
	if len(users) > 0 {
		return errors.New("用户名或邮箱已存在")
	}
	//没问题插入用户数据
	user := structs.User{
		Username: username,
		Password: password,
		Email:    email,
	}
	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
