package repositories

import (
	"github.com/harshit/buildgram/internal/models"
	"gorm.io/gorm"
)

// FollowRepository handles database operations for Follow model.
type FollowRepository struct {
	db *gorm.DB
}

// NewFollowRepository creates a new FollowRepository.
func NewFollowRepository(db *gorm.DB) *FollowRepository {
	return &FollowRepository{db: db}
}

// Create inserts a new follow relationship.
func (r *FollowRepository) Create(follow *models.Follow) error {
	return r.db.Create(follow).Error
}

// Delete removes a follow relationship (unfollow).
func (r *FollowRepository) Delete(followerID, followingID uint) error {
	return r.db.Where("follower_id = ? AND following_id = ?", followerID, followingID).Delete(&models.Follow{}).Error
}

// IsFollowing checks if a user is following another user.
func (r *FollowRepository) IsFollowing(followerID, followingID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.Follow{}).Where("follower_id = ? AND following_id = ?", followerID, followingID).Count(&count).Error
	return count > 0, err
}

// GetFollowers returns all followers for a user.
func (r *FollowRepository) GetFollowers(userID uint, offset, limit int) ([]models.User, error) {
	var users []models.User
	err := r.db.
		Joins("JOIN follows ON follows.follower_id = users.id").
		Where("follows.following_id = ?", userID).
		Offset(offset).
		Limit(limit).
		Find(&users).Error
	return users, err
}

// GetFollowing returns all users that a user is following.
func (r *FollowRepository) GetFollowing(userID uint, offset, limit int) ([]models.User, error) {
	var users []models.User
	err := r.db.
		Joins("JOIN follows ON follows.following_id = users.id").
		Where("follows.follower_id = ?", userID).
		Offset(offset).
		Limit(limit).
		Find(&users).Error
	return users, err
}
