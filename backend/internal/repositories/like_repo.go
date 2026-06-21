package repositories

import (
	"github.com/harshit/buildgram/internal/models"
	"gorm.io/gorm"
)

// LikeRepository handles database operations for Like model.
type LikeRepository struct {
	db *gorm.DB
}

// NewLikeRepository creates a new LikeRepository.
func NewLikeRepository(db *gorm.DB) *LikeRepository {
	return &LikeRepository{db: db}
}

// Create inserts a new like into the database.
func (r *LikeRepository) Create(like *models.Like) error {
	return r.db.Create(like).Error
}

// Delete removes a like (unlike).
func (r *LikeRepository) Delete(postID, userID uint) error {
	return r.db.Where("post_id = ? AND user_id = ?", postID, userID).Delete(&models.Like{}).Error
}

// IsLiked checks if a user has liked a specific post.
func (r *LikeRepository) IsLiked(postID, userID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.Like{}).Where("post_id = ? AND user_id = ?", postID, userID).Count(&count).Error
	return count > 0, err
}

// GetLikeCount returns the number of likes for a post.
func (r *LikeRepository) GetLikeCount(postID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Like{}).Where("post_id = ?", postID).Count(&count).Error
	return count, err
}
