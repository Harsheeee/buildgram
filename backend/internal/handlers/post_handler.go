package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/harshit/buildgram/internal/services"
)

// PostHandler handles post-related HTTP requests.
type PostHandler struct {
	postService *services.PostService
}

// NewPostHandler creates a new PostHandler.
func NewPostHandler(postService *services.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

// CreatePost creates a new post with an image.
// POST /api/posts
func (h *PostHandler) CreatePost(c *gin.Context) {
	currentUserID, _ := c.Get("userID")

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "image file is required"})
		return
	}
	defer file.Close()

	caption := c.PostForm("caption")

	response, err := h.postService.CreatePost(currentUserID.(uint), caption, file, header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetFeed retrieves the authenticated user's feed.
// GET /api/posts/feed?page=1&limit=10
func (h *PostHandler) GetFeed(c *gin.Context) {
	currentUserID, _ := c.Get("userID")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 50 {
		limit = 10
	}

	posts, err := h.postService.GetFeed(currentUserID.(uint), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch feed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts, "page": page, "limit": limit})
}

// GetUserPosts retrieves all posts by a specific user.
// GET /api/posts/user/:id?page=1&limit=10
func (h *PostHandler) GetUserPosts(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	currentUserID, _ := c.Get("userID")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "12"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 50 {
		limit = 12
	}

	posts, err := h.postService.GetUserPosts(uint(userID), currentUserID.(uint), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts, "page": page, "limit": limit})
}

// GetPost retrieves a single post by ID.
// GET /api/posts/:id
func (h *PostHandler) GetPost(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	currentUserID, _ := c.Get("userID")

	post, err := h.postService.GetPost(uint(postID), currentUserID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

// DeletePost deletes a post.
// DELETE /api/posts/:id
func (h *PostHandler) DeletePost(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	currentUserID, _ := c.Get("userID")

	if err := h.postService.DeletePost(uint(postID), currentUserID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post deleted"})
}

// GetExplorePosts retrieves posts for the explore page.
// GET /api/posts/explore?page=1&limit=20
func (h *PostHandler) GetExplorePosts(c *gin.Context) {
	currentUserID, _ := c.Get("userID")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 50 {
		limit = 20
	}

	posts, err := h.postService.GetExplorePosts(currentUserID.(uint), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch explore posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts, "page": page, "limit": limit})
}
