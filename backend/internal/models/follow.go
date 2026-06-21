package models

import (
	"time"
)

// Follow represents a follow relationship between two users.
type Follow struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	FollowerID  uint      `gorm:"not null;uniqueIndex:idx_follower_following" json:"follower_id"`
	FollowingID uint      `gorm:"not null;uniqueIndex:idx_follower_following" json:"following_id"`
	CreatedAt   time.Time `json:"created_at"`

	// Relationships
	Follower  User `gorm:"foreignKey:FollowerID" json:"follower,omitempty"`
	Following User `gorm:"foreignKey:FollowingID" json:"following,omitempty"`
}
