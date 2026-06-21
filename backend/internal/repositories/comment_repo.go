package repositories

import (
	"github.com/harshit/buildgram/internal/models"
	"gorm.io/gorm"
)

// CommentRepository handles database operations for Comment model.
type CommentRepository struct {
	db *gorm.DB
}

// NewCommentRepository creates a new CommentRepository.
func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

// Create inserts a new comment into the database.
func (r *CommentRepository) Create(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

// GetByPostID retrieves all comments for a specific post.
func (r *CommentRepository) GetByPostID(postID uint, offset, limit int) ([]models.Comment, error) {
	var comments []models.Comment
	err := r.db.
		Preload("User").
		Where("post_id = ?", postID).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&comments).Error
	return comments, err
}

// Delete removes a comment.
func (r *CommentRepository) Delete(id, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Comment{}).Error
}

// GetCommentCount returns the number of comments for a post.
func (r *CommentRepository) GetCommentCount(postID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Comment{}).Where("post_id = ?", postID).Count(&count).Error
	return count, err
}
