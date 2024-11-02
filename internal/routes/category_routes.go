package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nicoalexanderdev/api-portafolio/internal/controllers"
	"github.com/nicoalexanderdev/api-portafolio/internal/repositories"
	"github.com/nicoalexanderdev/api-portafolio/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupCategoryRoutes(router *gin.Engine, db *mongo.Database) {
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryController := controllers.NewCategoryController(categoryService)

	categoryRoutes := router.Group("/api/v1/categories")
	{
		categoryRoutes.POST("", categoryController.CreateCategory)
		categoryRoutes.GET("", categoryController.GetAllCategories)
		categoryRoutes.GET("/:id", categoryController.GetCategoryByID)
		categoryRoutes.PUT("/:id", categoryController.UpdateCategory)
		categoryRoutes.DELETE("/:id", categoryController.DeleteCategory)
	}
}
