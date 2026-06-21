package repositories

import (
	"github.com/harshit/buildgram/internal/models"
	"gorm.io/gorm"
)

// UserRepository handles database operations for User model.
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create inserts a new user into the database.
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// FindByID retrieves a user by their ID.
func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

// FindByEmail retrieves a user by email.
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

// FindByUsername retrieves a user by username.
func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

// Update updates a user's information.
func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

// GetFollowerCount returns the number of followers for a user.
func (r *UserRepository) GetFollowerCount(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Follow{}).Where("following_id = ?", userID).Count(&count).Error
	return count, err
}

// GetFollowingCount returns the number of users a user is following.
func (r *UserRepository) GetFollowingCount(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Follow{}).Where("follower_id = ?", userID).Count(&count).Error
	return count, err
}

// GetPostCount returns the number of posts for a user.
func (r *UserRepository) GetPostCount(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Post{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}

// SearchUsers searches for users by username (partial match).
func (r *UserRepository) SearchUsers(query string, limit int) ([]models.User, error) {
	var users []models.User
	err := r.db.Where("username ILIKE ?", "%"+query+"%").Limit(limit).Find(&users).Error
	return users, err
}
