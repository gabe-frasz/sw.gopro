package usecases

import (
	"github.com/gabe-frasz/gopro/tasks/internal/app/entity"
	"github.com/gabe-frasz/gopro/tasks/internal/app/repository"
)

type ListAllTasks struct {
	repository repository.TaskRepository
}

func NewListAllTasks(repository repository.TaskRepository) *ListAllTasks {
	return &ListAllTasks{
		repository: repository,
	}
}

func (l *ListAllTasks) Execute() ([]*entity.Task, error) {
	return l.repository.GetAll()
}
