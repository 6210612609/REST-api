package service

import "example.com/entity"

type ProblemService interface {
	Commit(entity.Problem) entity.Problem
	FindAllProblem() []entity.Problem
}

type problemService struct {
	problems []entity.Problem
}

func NewProblem() ProblemService {
	return &problemService{
		problems: []entity.Problem{},
	}
}

func (service *problemService) Commit(problem entity.Problem) entity.Problem {
	service.problems = append(service.problems, problem)
	return problem
}

func (service *problemService) FindAllProblem() []entity.Problem {
	return service.problems
}
