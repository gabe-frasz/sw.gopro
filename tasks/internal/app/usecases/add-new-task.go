package usecases

import (
	"github.com/gabe-frasz/gopro/tasks/internal/app/entity"
	"github.com/gabe-frasz/gopro/tasks/internal/app/repository"
)

type AddNewTask struct {
	description string
	repository  repository.TaskRepository
}

func NewAddNewTask(description string, repository repository.TaskRepository) *AddNewTask {
	return &AddNewTask{
		description: description,
		repository:  repository,
	}
}

func (a *AddNewTask) Execute() error {
	task := *entity.NewTask(0, a.description, false)
	return a.repository.Add(&task)
}
