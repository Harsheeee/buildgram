package services

import (
	"errors"

	"github.com/harshit/buildgram/internal/models"
	"github.com/harshit/buildgram/internal/repositories"
)

// InteractionService handles likes, comments, and follows business logic.
type InteractionService struct {
	likeRepo    *repositories.LikeRepository
	commentRepo *repositories.CommentRepository
	followRepo  *repositories.FollowRepository
	userRepo    *repositories.UserRepository
}

// NewInteractionService creates a new InteractionService.
func NewInteractionService(
	likeRepo *repositories.LikeRepository,
	commentRepo *repositories.CommentRepository,
	followRepo *repositories.FollowRepository,
	userRepo *repositories.UserRepository,
) *InteractionService {
	return &InteractionService{
		likeRepo:    likeRepo,
		commentRepo: commentRepo,
		followRepo:  followRepo,
		userRepo:    userRepo,
	}
}

// ToggleLike toggles a like on a post. Returns true if liked, false if unliked.
func (s *InteractionService) ToggleLike(postID, userID uint) (bool, int64, error) {
	isLiked, err := s.likeRepo.IsLiked(postID, userID)
	if err != nil {
		return false, 0, err
	}

	if isLiked {
		if err := s.likeRepo.Delete(postID, userID); err != nil {
			return false, 0, errors.New("failed to unlike post")
		}
	} else {
		like := &models.Like{
			PostID: postID,
			UserID: userID,
		}
		if err := s.likeRepo.Create(like); err != nil {
			return false, 0, errors.New("failed to like post")
		}
	}

	count, _ := s.likeRepo.GetLikeCount(postID)
	return !isLiked, count, nil
}

// AddComment adds a comment to a post.
func (s *InteractionService) AddComment(postID, userID uint, content string) (*CommentResponse, error) {
	if content == "" {
		return nil, errors.New("comment content cannot be empty")
	}

	comment := &models.Comment{
		PostID:  postID,
		UserID:  userID,
		Content: content,
	}

	if err := s.commentRepo.Create(comment); err != nil {
		return nil, errors.New("failed to add comment")
	}

	user, _ := s.userRepo.FindByID(userID)

	return &CommentResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		User: UserResponse{
			ID:                user.ID,
			Username:          user.Username,
			ProfilePictureURL: user.ProfilePictureURL,
		},
	}, nil
}

// GetComments retrieves comments for a post.
func (s *InteractionService) GetComments(postID uint, page, limit int) ([]CommentResponse, error) {
	offset := (page - 1) * limit
	comments, err := s.commentRepo.GetByPostID(postID, offset, limit)
	if err != nil {
		return nil, err
	}

	var responses []CommentResponse
	for _, c := range comments {
		responses = append(responses, CommentResponse{
			ID:        c.ID,
			Content:   c.Content,
			CreatedAt: c.CreatedAt,
			User: UserResponse{
				ID:                c.User.ID,
				Username:          c.User.Username,
				ProfilePictureURL: c.User.ProfilePictureURL,
			},
		})
	}
	return responses, nil
}

// DeleteComment deletes a comment.
func (s *InteractionService) DeleteComment(commentID, userID uint) error {
	return s.commentRepo.Delete(commentID, userID)
}

// ToggleFollow toggles a follow relationship. Returns true if now following, false if unfollowed.
func (s *InteractionService) ToggleFollow(followerID, followingID uint) (bool, error) {
	if followerID == followingID {
		return false, errors.New("cannot follow yourself")
	}

	// Verify the target user exists
	_, err := s.userRepo.FindByID(followingID)
	if err != nil {
		return false, errors.New("user not found")
	}

	isFollowing, err := s.followRepo.IsFollowing(followerID, followingID)
	if err != nil {
		return false, err
	}

	if isFollowing {
		if err := s.followRepo.Delete(followerID, followingID); err != nil {
			return false, errors.New("failed to unfollow")
		}
		return false, nil
	}

	follow := &models.Follow{
		FollowerID:  followerID,
		FollowingID: followingID,
	}
	if err := s.followRepo.Create(follow); err != nil {
		return false, errors.New("failed to follow")
	}
	return true, nil
}

// GetFollowers returns followers for a user.
func (s *InteractionService) GetFollowers(userID uint, page, limit int) ([]UserResponse, error) {
	offset := (page - 1) * limit
	users, err := s.followRepo.GetFollowers(userID, offset, limit)
	if err != nil {
		return nil, err
	}

	var responses []UserResponse
	for _, u := range users {
		responses = append(responses, UserResponse{
			ID:                u.ID,
			Username:          u.Username,
			FullName:          u.FullName,
			ProfilePictureURL: u.ProfilePictureURL,
		})
	}
	return responses, nil
}

// GetFollowing returns users that a user is following.
func (s *InteractionService) GetFollowing(userID uint, page, limit int) ([]UserResponse, error) {
	offset := (page - 1) * limit
	users, err := s.followRepo.GetFollowing(userID, offset, limit)
	if err != nil {
		return nil, err
	}

	var responses []UserResponse
	for _, u := range users {
		responses = append(responses, UserResponse{
			ID:                u.ID,
			Username:          u.Username,
			FullName:          u.FullName,
			ProfilePictureURL: u.ProfilePictureURL,
		})
	}
	return responses, nil
}
