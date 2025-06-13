package structs

import "time"

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
	UserID    uint   `gorm:"not null"`
	PostID    uint   `gorm:"not null"`
	IsDel     uint8  `gorm:"default:0;not null"`
	CreatedAt time.Time
	User      User
	Post      Post
}
