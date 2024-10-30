package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nicoalexanderdev/api-portafolio/controller"
	"github.com/nicoalexanderdev/api-portafolio/service"
)

var (
	projectService    service.ProjectService       = service.New()
	projectController controller.ProjectController = controller.New(projectService)
)

func main() {
	server := gin.Default()

	server.GET("/projects", func(ctx *gin.Context) {
		ctx.JSON(200, projectController.FindAll())
	})

	server.POST("/projects", func(ctx *gin.Context) {
		ctx.JSON(200, projectController.Save(ctx))
	})

	server.Run(":8080")
}
