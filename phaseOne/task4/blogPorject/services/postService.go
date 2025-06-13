package services

import (
	"errors"
	"gotask/phaseOne/task4/blogPorject/config"
	"gotask/phaseOne/task4/blogPorject/structs"
)

func CreatePost(post structs.Post) error {
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
	result := config.DB.Preload("Comments").Preload("User").First(post, postId)
	return result.Error
}
