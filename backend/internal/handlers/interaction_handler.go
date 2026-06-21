package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/harshit/buildgram/internal/services"
)

// InteractionHandler handles likes, comments, and follows HTTP requests.
type InteractionHandler struct {
	interactionService *services.InteractionService
}

// NewInteractionHandler creates a new InteractionHandler.
func NewInteractionHandler(interactionService *services.InteractionService) *InteractionHandler {
	return &InteractionHandler{interactionService: interactionService}
}

// ToggleLike toggles a like on a post.
// POST /api/posts/:id/like
func (h *InteractionHandler) ToggleLike(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	currentUserID, _ := c.Get("userID")

	isLiked, likeCount, err := h.interactionService.ToggleLike(uint(postID), currentUserID.(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"is_liked":   isLiked,
		"like_count": likeCount,
	})
}

// AddComment adds a comment to a post.
// POST /api/posts/:id/comments
func (h *InteractionHandler) AddComment(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	currentUserID, _ := c.Get("userID")

	var input struct {
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "comment content is required"})
		return
	}

	comment, err := h.interactionService.AddComment(uint(postID), currentUserID.(uint), input.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// GetComments retrieves comments for a post.
// GET /api/posts/:id/comments?page=1&limit=20
func (h *InteractionHandler) GetComments(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	comments, err := h.interactionService.GetComments(uint(postID), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch comments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comments": comments, "page": page, "limit": limit})
}

// DeleteComment deletes a comment.
// DELETE /api/comments/:id
func (h *InteractionHandler) DeleteComment(c *gin.Context) {
	commentID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid comment ID"})
		return
	}

	currentUserID, _ := c.Get("userID")

	if err := h.interactionService.DeleteComment(uint(commentID), currentUserID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "comment deleted"})
}

// ToggleFollow toggles a follow relationship.
// POST /api/users/:id/follow
func (h *InteractionHandler) ToggleFollow(c *gin.Context) {
	followingID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	currentUserID, _ := c.Get("userID")

	isFollowing, err := h.interactionService.ToggleFollow(currentUserID.(uint), uint(followingID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"is_following": isFollowing})
}

// GetFollowers returns followers of a user.
// GET /api/users/:id/followers?page=1&limit=20
func (h *InteractionHandler) GetFollowers(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	users, err := h.interactionService.GetFollowers(uint(userID), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch followers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users, "page": page, "limit": limit})
}

// GetFollowing returns users that a user is following.
// GET /api/users/:id/following?page=1&limit=20
func (h *InteractionHandler) GetFollowing(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	users, err := h.interactionService.GetFollowing(uint(userID), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch following"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users, "page": page, "limit": limit})
}
