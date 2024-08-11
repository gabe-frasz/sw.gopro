package usecases

import (
	"github.com/gabe-frasz/gopro/tasks/internal/app/entity"
	"github.com/gabe-frasz/gopro/tasks/internal/app/repository"
)

type ListUncompletedTasks struct {
	repository repository.TaskRepository
}

func NewListUncompletedTasks(repository repository.TaskRepository) *ListUncompletedTasks {
	return &ListUncompletedTasks{
		repository: repository,
	}
}

func (l *ListUncompletedTasks) Execute() ([]*entity.Task, error) {
	return l.repository.GetUncompleted()
}
