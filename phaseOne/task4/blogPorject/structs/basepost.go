package structs

import "time"

type Post struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Content   string `gorm:"type:text;not null"`
	UserID    uint   `gorm:"not null"`
	IsDel     uint8  `gorm:"default:0;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
	Comments  []Comment
}
