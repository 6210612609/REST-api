package controller

import (
	"example.com/entity"
	"example.com/service"
	"github.com/gin-gonic/gin"
)

type ProblemController interface {
	FindAllProblem() []entity.Problem
	Commit(ctx *gin.Context) entity.Problem
}

type controllerProblem struct {
	servicepro service.ProblemService
}

func NewProblem(service service.ProblemService) ProblemController {
	return &controllerProblem{
		servicepro: service,
	}
}

func (c *controllerProblem) FindAllProblem() []entity.Problem {
	return c.servicepro.FindAllProblem()
}

func (c *controllerProblem) Commit(ctx *gin.Context) entity.Problem {
	var problem entity.Problem
	ctx.ShouldBind(&problem)
	c.servicepro.Commit(problem)
	return problem
}
