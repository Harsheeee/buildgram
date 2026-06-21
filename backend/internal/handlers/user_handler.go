package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/harshit/buildgram/internal/services"
)

// UserHandler handles user-related HTTP requests.
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetProfile retrieves a user's profile.
// GET /api/users/:id
func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	currentUserID, _ := c.Get("userID")

	profile, err := h.userService.GetProfile(uint(userID), currentUserID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

// GetProfileByUsername retrieves a user's profile by username.
// GET /api/users/username/:username
func (h *UserHandler) GetProfileByUsername(c *gin.Context) {
	// This is handled via search + profile fetch on the frontend
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// UpdateProfile updates the current user's profile.
// PUT /api/users/profile
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	currentUserID, _ := c.Get("userID")

	var input services.UpdateProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.userService.UpdateProfile(currentUserID.(uint), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// UploadProfilePicture uploads a profile picture for the current user.
// POST /api/users/profile/picture
func (h *UserHandler) UploadProfilePicture(c *gin.Context) {
	currentUserID, _ := c.Get("userID")

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "image file is required"})
		return
	}
	defer file.Close()

	response, err := h.userService.UploadProfilePicture(currentUserID.(uint), file, header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// SearchUsers searches for users by username.
// GET /api/users/search?q=query
func (h *UserHandler) SearchUsers(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "search query is required"})
		return
	}

	users, err := h.userService.SearchUsers(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to search users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
