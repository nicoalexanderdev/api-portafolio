package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicoalexanderdev/api-portafolio/internal/models"
	"github.com/nicoalexanderdev/api-portafolio/internal/services"
)

type ProjectController struct {
	service services.ProjectService
}

func NewProjectController(service services.ProjectService) *ProjectController {
	return &ProjectController{
		service: service,
	}
}

func (c *ProjectController) CreateProject(ctx *gin.Context) {
	var project models.Project
	if err := ctx.ShouldBindJSON(&project); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateProject(ctx.Request.Context(), &project); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, project)
}

func (c *ProjectController) GetAllProjects(ctx *gin.Context) {
	projects, err := c.service.GetAllProjects(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, projects)
}

func (c *ProjectController) GetProjectByID(ctx *gin.Context) {
	id := ctx.Param("id")
	project, err := c.service.GetProjectByID(ctx.Request.Context(), id)
	if err != nil {
		switch err {
		case services.ErrProjectNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		case services.ErrInvalidID:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, project)
}

func (c *ProjectController) UpdateProject(ctx *gin.Context) {
	id := ctx.Param("id")
	var project models.Project
	if err := ctx.ShouldBindJSON(&project); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.UpdateProject(ctx.Request.Context(), id, &project); err != nil {
		switch err {
		case services.ErrProjectNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		case services.ErrInvalidID:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Project updated successfully"})
}

func (c *ProjectController) DeleteProject(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeleteProject(ctx.Request.Context(), id); err != nil {
		switch err {
		case services.ErrProjectNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		case services.ErrInvalidID:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}
