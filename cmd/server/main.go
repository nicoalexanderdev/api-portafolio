package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nicoalexanderdev/api-portafolio/config"
	"github.com/nicoalexanderdev/api-portafolio/internal/routes"
)

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

	// Configuración simple de CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:4200", "https://portafolio-evjm.onrender.com"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	// Setup routes
	routes.SetupProjectRoutes(router, mongoDatabase)
	routes.SetupCategoryRoutes(router, mongoDatabase)
	routes.SetupBlogRoutes(router, mongoDatabase)

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
