package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nicoalexanderdev/api-portafolio/entity"
	"github.com/nicoalexanderdev/api-portafolio/service"
)

type ProjectController interface {
	FindAll() []entity.Project
	Save(ctx *gin.Context) error
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
func (c *controller) Save(ctx *gin.Context) error {
	var project entity.Project
	err := ctx.ShouldBindJSON(&project)
	if err != nil {
		return err
	}
	c.service.Save(project)
	return nil
}
