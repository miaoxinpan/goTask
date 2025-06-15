package services

import (
	"errors"
	"gotask/phaseOne/task4/blogPorject/config"
	"gotask/phaseOne/task4/blogPorject/structs"
	"gotask/phaseOne/task4/blogPorject/utils"
)

func CreatePost(post structs.Post) error {
	utils.LogBusiness("CreatePost")
	// 业务字段校验
	if post.Title == "" || post.Content == "" {
		return errors.New("标题或内容不能为空！")

	}
	//接下来保存文章信息
	//save  更新 或者  新增 都可以   全量的
	//Create  只做插入
	if err := config.DB.Create(&post).Error; err != nil {
		return err
	}
	return nil
}
func GetAllPost(page int, pageSize int, posts *[]structs.Post, total *int64) error {
	utils.LogBusiness("GetAllPost")
	offset := (page - 1) * pageSize
	db := config.DB.Model(&structs.Post{})
	if err := db.Count(total).Error; err != nil {
		return err
	}
	if err := db.Preload("User").Order("created_at desc").Limit(pageSize).Offset(offset).Find(posts).Error; err != nil {
		return err
	}
	return nil
}

//	func GetPostForId(postId uint, post *structs.Post) error {
//		err := config.DB.Preload("Comments").Model(&structs.Post{}).Where("id = ?", postId).First(post)
//		if err != nil {
//			return err.Error
//		}
//		return nil
//	}
func GetPostForId(postId uint, post *structs.Post) error {
	utils.LogBusiness("GetPostForId")
	result := config.DB.Preload("Comments").Preload("User").First(&post, postId)
	return result.Error
}

func UpdatePostForAuthor(postId uint, content string, opType string) error {
	utils.LogBusiness("UpdatePostForAuthor")
	if "D" == opType {
		return config.DB.Model(&structs.Post{}).Where("id = ?", postId).Update("is_del", 1).Error
	} else if "U" == opType {
		return config.DB.Model(&structs.Post{}).Where("id = ?", postId).Update("content", content).Error
	}
	return errors.New("不支持的操作类型")
}
func GetPostForUserId(userid uint, post *[]structs.Post) error {
	utils.LogBusiness("GetPostForUserId")
	err := config.DB.Preload("User").Where("user_id=?", userid).Find(&post).Error
	return err
}
