package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/harshit/buildgram/internal/config"
	"github.com/harshit/buildgram/internal/handlers"
	"github.com/harshit/buildgram/internal/middlewares"
	"github.com/harshit/buildgram/internal/models"
	"github.com/harshit/buildgram/internal/repositories"
	"github.com/harshit/buildgram/internal/routes"
	"github.com/harshit/buildgram/internal/services"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db := config.InitDB(cfg)

	// Run migrations
	models.AutoMigrate(db)

	// Create uploads directory
	if err := os.MkdirAll(cfg.UploadDir+"/posts", os.ModePerm); err != nil {
		log.Fatalf("Failed to create uploads directory: %v", err)
	}
	if err := os.MkdirAll(cfg.UploadDir+"/profiles", os.ModePerm); err != nil {
		log.Fatalf("Failed to create profiles directory: %v", err)
	}

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	postRepo := repositories.NewPostRepository(db)
	commentRepo := repositories.NewCommentRepository(db)
	likeRepo := repositories.NewLikeRepository(db)
	followRepo := repositories.NewFollowRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	userService := services.NewUserService(userRepo, followRepo, cfg.UploadDir)
	postService := services.NewPostService(postRepo, likeRepo, commentRepo, cfg.UploadDir)
	interactionService := services.NewInteractionService(likeRepo, commentRepo, followRepo, userRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	postHandler := handlers.NewPostHandler(postService)
	interactionHandler := handlers.NewInteractionHandler(interactionService)

	// Setup Gin router
	router := gin.Default()

	// Middleware
	router.Use(middlewares.CORSMiddleware())

	// Serve uploaded files statically
	router.Static("/uploads", cfg.UploadDir)

	// Setup routes
	routes.SetupRoutes(router, cfg.JWTSecret, authHandler, userHandler, postHandler, interactionHandler)

	// Start server
	log.Printf("🚀 BuildGram server starting on port %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
