package controller

import (
	"example.com/entity"
	"example.com/service"
	"github.com/gin-gonic/gin"
)

type AssignmentController interface {
	FindAll() []entity.Assignment
	Save(ctx *gin.Context) entity.Assignment
}

type controller struct {
	service service.AssignmentService
}

func New(service service.AssignmentService) AssignmentController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Assignment {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Assignment {
	var assignment entity.Assignment
	ctx.ShouldBind(&assignment)
	c.service.Save(assignment)
	return assignment
}
