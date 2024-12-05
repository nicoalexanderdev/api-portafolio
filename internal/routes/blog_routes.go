package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nicoalexanderdev/api-portafolio/internal/controllers"
	"github.com/nicoalexanderdev/api-portafolio/internal/repositories"
	"github.com/nicoalexanderdev/api-portafolio/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupBlogRoutes(router *gin.Engine, db *mongo.Database) {
	blogRepo := repositories.NewBlogRepository(db)
	blogService := services.NewBlogService(blogRepo)
	blogController := controllers.NewBlogController(blogService)

	blogRoutes := router.Group("/api/v1/blogs")
	{
		blogRoutes.POST("", blogController.CreateBlog)
		blogRoutes.GET("", blogController.GetAllBlogs)
		blogRoutes.GET("/:id", blogController.GetBlogByID)
		blogRoutes.GET("/category/:categoryId", blogController.GetBlogsByCategory)
		blogRoutes.PUT("/:id", blogController.UpdateBlog)
		blogRoutes.DELETE("/:id", blogController.DeleteBlog)
	}
}
