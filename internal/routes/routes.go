package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nicoalexanderdev/api-portafolio/internal/controllers"
	"github.com/nicoalexanderdev/api-portafolio/internal/repositories"
	"github.com/nicoalexanderdev/api-portafolio/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupProjectRoutes(router *gin.Engine, db *mongo.Database) {
	projectRepo := repositories.NewProjectRepository(db)
	projectService := services.NewProjectService(projectRepo)
	projectController := controllers.NewProjectController(projectService)

	projectRoutes := router.Group("/api/v1/projects")
	{
		projectRoutes.POST("", projectController.CreateProject)
		projectRoutes.GET("", projectController.GetAllProjects)
		projectRoutes.GET("/:id", projectController.GetProjectByID)
		projectRoutes.PUT("/:id", projectController.UpdateProject)
		projectRoutes.DELETE("/:id", projectController.DeleteProject)
	}

	// Protected routes
	// protected := r.Group("/api/v1")
	// protected.Use(middleware.AuthMiddleware())
	// {
	// 	protected.GET("/profile", controllers.GetProfile)
	// 	protected.PUT("/profile", controllers.UpdateProfile)
	// }
}
