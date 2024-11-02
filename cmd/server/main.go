package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicoalexanderdev/api-portafolio/config"
	"github.com/nicoalexanderdev/api-portafolio/internal/routes"
)

// var (
// 	projectService    service.ProjectService       = service.New()
// 	projectController controller.ProjectController = controller.New(projectService)
// )

func main() {
	// Load configuration
	cfg := config.GetConfig()

	// Set Gin mode
	gin.SetMode(cfg.Server.Mode)

	// Connect to MongoDB
	mongoClient, err := config.ConnectMongoDB(cfg)
	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}
	defer mongoClient.Disconnect(context.Background())

	// Get specific database from client
	mongoDatabase := mongoClient.Database(cfg.Database.Name)

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupProjectRoutes(router, mongoDatabase)
	routes.SetupCategoryRoutes(router, mongoDatabase)

	// Start server
	server := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: router,
	}

	log.Printf("Starting server on port %s", cfg.Server.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server startup failed: %v", err)
	}
}
