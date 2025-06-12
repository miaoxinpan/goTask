package code

import "gorm.io/gorm"

/*
进阶gorm
	题目1：模型定义
		假设你要开发一个博客系统，有以下几个实体：
		User （用户）、 Post （文章）、 Comment （评论）。
	要求 ：
		使用Gorm定义 User 、 Post 和 Comment 模型，
		其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
		Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
		编写Go代码，使用Gorm创建这些模型对应的数据库表。
	题目2：关联查询
		基于上述博客系统的模型定义。
	要求 ：
		编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
		编写Go代码，使用Gorm查询评论数量最多的文章信息。
	题目3：钩子函数
		继续使用博客系统的模型。
	要求 ：
		为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。

		为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
type User struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Posts     []Post `gorm:"foreignKey:UserID"`
	PostCount int
}

type Post struct {
	ID         int `gorm:"primaryKey"`
	Title      string
	Content    string
	UserID     int
	Comments   []Comment `gorm:"foreignKey:PostID"`
	PostStatus string
}

type Comment struct {
	ID      int `gorm:"primaryKey"`
	Content string
	PostID  int
	UserID  int
}

//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
func QueryPostDetailInfo(db *gorm.DB, userid int) []Post {
	var posts []Post
	db.Preload("Comments").Where("user_id = ? ", userid).Find(&posts)
	return posts
}

//编写Go代码，使用Gorm查询评论数量最多的文章信息。
func QueryMaxCommentForPost(db *gorm.DB) Post {
	var post Post
	result := make(map[string]interface{})
	db.Raw("SELECT post_id, COUNT(post_id) as cnt FROM comments GROUP BY post_id ORDER BY cnt DESC LIMIT 1").Scan(&result)
	db.Where("id = ? ", result["post_id"]).First(&post)
	return post
}

/*
	为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段
	这个是文章创建以后 更新用户信息  所以是AfterCreate

*/
func PublishPost(db *gorm.DB, post *Post) error {
	return db.Create(post).Error
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	var count int64
	tx.Raw("select count(1) from posts where user_id = ?", p.UserID).Scan(&count)
	tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", count)
	return
}

/*
	为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，
	如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	删除 可用的hook  是BeforeDelete AfterDelete
	这个是删除评论之后  所以也是AfterDelete

	延伸出另外一个钩子  就是创建评论后
	看这个文章的评论状态 是不是 无评论
	是的话  得给他清空这个字段 或者改成 有评论


*/
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	//根据传来的c 查一下这个post还有没有评论
	var count int64
	tx.Raw("select count(1) from comments where   post_id = ?", c.PostID).Scan(&count)
	if count == 0 { //如果没有评论了
		tx.Model(&Post{}).Where("id = ?", c.PostID).Update("post_status", "无评论")
	}
	return
}
func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	//直接查询这条post的这个字段 是不是 无评论吧
	//新建一个post对象用来接收
	var post Post
	tx.Where("id = ? ", c.PostID).First(&post)
	// 如果状态为"无评论"，则改为"有评论"
	if post.PostStatus == "无评论" {
		tx.Model(&post).Update("post_status", "有评论")
	}
	return
}

//发表评论
func MakeComment(db *gorm.DB, comment *Comment) error {
	return db.Create(comment).Error
}

//删除评论
func DelectComment(db *gorm.DB, id int) {
	db.Delete(&Comment{}, id)
}
