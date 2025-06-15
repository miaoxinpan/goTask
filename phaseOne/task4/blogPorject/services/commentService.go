package services

import (
	"gotask/phaseOne/task4/blogPorject/config"
	"gotask/phaseOne/task4/blogPorject/structs"
	"gotask/phaseOne/task4/blogPorject/utils"
)

func CreateComment(postId uint, content string, userId uint) error {
	utils.LogBusiness("CreateComment")
	comment := structs.Comment{
		PostID:  postId,
		Content: content,
		UserID:  userId,
	}
	if err := config.DB.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

func GetCommentForUserId(userid uint, comments *[]structs.Comment) error {
	utils.LogBusiness("GetCommentForUserId")
	err := config.DB.Where("user_id=?", userid).Order("created_at asc").Find(&comments)
	return err.Error
}

func GetCommentForPostId(postId uint, comments *[]structs.Comment) error {
	utils.LogBusiness("GetCommentForPostId")
	result := config.DB.Where("post_id=?", postId).Order("created_at asc").Find(&comments)
	return result.Error
}
