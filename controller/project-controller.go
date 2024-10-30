package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nicoalexanderdev/api-portafolio/entity"
	"github.com/nicoalexanderdev/api-portafolio/service"
)

type ProjectController interface {
	FindAll() []entity.Project
	Save(ctx *gin.Context) entity.Project
}

type controller struct {
	service service.ProjectService
}

func New(service service.ProjectService) ProjectController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Project {
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) entity.Project {
	var project entity.Project
	ctx.BindJSON(&project)
	c.service.Save(project)
	return project
}
