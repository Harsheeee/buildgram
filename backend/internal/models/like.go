package models

import (
	"time"
)

// Like represents a like on a post by a user.
type Like struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PostID    uint      `gorm:"not null;uniqueIndex:idx_user_post_like" json:"post_id"`
	UserID    uint      `gorm:"not null;uniqueIndex:idx_user_post_like" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`

	// Relationships
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Post Post `gorm:"foreignKey:PostID" json:"-"`
}
