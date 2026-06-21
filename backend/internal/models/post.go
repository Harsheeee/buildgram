package models

import (
	"time"

	"gorm.io/gorm"
)

// Post represents a user's image post.
type Post struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"not null;index" json:"user_id"`
	ImageURL  string         `gorm:"size:500;not null" json:"image_url"`
	Caption   string         `gorm:"size:2200" json:"caption"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	User     User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Comments []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty"`
	Likes    []Like    `gorm:"foreignKey:PostID" json:"likes,omitempty"`
}
