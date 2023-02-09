package service

import "example.com/entity"

type AssignmentService interface {
	Save(entity.Assignment) entity.Assignment
	FindAll() []entity.Assignment
}

type assignmentService struct {
	assignments []entity.Assignment
}

func New() AssignmentService {

	return &assignmentService{
		assignments: []entity.Assignment{},
	}
}

func (service *assignmentService) Save(assignment entity.Assignment) entity.Assignment {
	service.assignments = append(service.assignments, assignment)
	return assignment
}

func (service *assignmentService) FindAll() []entity.Assignment {
	return service.assignments
}
