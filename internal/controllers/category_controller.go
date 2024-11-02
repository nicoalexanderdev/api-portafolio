package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicoalexanderdev/api-portafolio/internal/models"
	"github.com/nicoalexanderdev/api-portafolio/internal/services"
)

// conectamos con services
type CategoryController struct {
	service services.CategoryService
}

// constructor
func NewCategoryController(service services.CategoryService) *CategoryController {
	return &CategoryController{
		service: service,
	}
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var category models.Category

	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := c.service.CreateCategory(ctx.Request.Context(), &category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, category)
}

func (c *CategoryController) GetAllCategories(ctx *gin.Context) {
	categories, err := c.service.GetAllCategories(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

func (c *CategoryController) GetCategoryByID(ctx *gin.Context) {
	id := ctx.Param("id")
	category, err := c.service.GetCategoryByID(ctx.Request.Context(), id)
	if err != nil {
		switch err {
		case services.ErrCategoryNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"Error": "Category not found"})
		case services.ErrInvalidID:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, category)
}

func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	var category models.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.UpdateCategory(ctx.Request.Context(), id, &category); err != nil {
		switch err {
		case services.ErrCategoryNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		case services.ErrInvalidID:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeleteCategory(ctx.Request.Context(), id); err != nil {
		switch err {
		case services.ErrCategoryNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		case services.ErrInvalidID:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
