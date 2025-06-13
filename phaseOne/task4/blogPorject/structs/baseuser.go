package structs

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	IsDel     uint8  `gorm:"default:0;not null"` // 0=未删除，1=已删除
	CreatedAt time.Time
	UpdatedAt time.Time
	Posts     []Post
	Comments  []Comment
}
