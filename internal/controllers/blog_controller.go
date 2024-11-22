package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicoalexanderdev/api-portafolio/internal/models"
	"github.com/nicoalexanderdev/api-portafolio/internal/services"
)

type BlogController struct {
	service services.BlogService
}

func NewBlogController(service services.BlogService) *BlogController {
	return &BlogController{
		service: service,
	}
}

func (c *BlogController) CreateBlog(ctx *gin.Context) {
	var blog models.Blog

	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := c.service.CreateBlog(ctx.Request.Context(), &blog); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, blog)
}

func (c *BlogController) GetAllBlogs(ctx *gin.Context) {
	blogs, err := c.service.GetAllBlogs(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}

func (c *BlogController) GetBlogByID(ctx *gin.Context) {
	id := ctx.Param("id")
	blog, err := c.service.GetBlogByID(ctx.Request.Context(), id)
	if err != nil {
		switch err {
		case services.ErrBlogNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"Error": "Blog Not Found"})
		case services.ErrInvalidID:
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Blog Id"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

func (c *BlogController) UpdateBlog(ctx *gin.Context) {
	id := ctx.Param("id")
	var blog models.Blog
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.UpdateBlog(ctx.Request.Context(), id, &blog); err != nil {
		switch err {
		case services.ErrBlogNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"Error": "Blog not found"})
		case services.ErrInvalidID:
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Blog ID"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Blog updated successfully"})
}

func (c *BlogController) DeleteBlog(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeleteBlog(ctx.Request.Context(), id); err != nil {
		switch err {
		case services.ErrBlogNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"Error": "Blog not found"})
		case services.ErrInvalidID:
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Blog ID"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}
