package usecases

import (
	"github.com/gabe-frasz/gopro/tasks/internal/app/repository"
)

type DeleteTask struct {
	id         int
	repository repository.TaskRepository
}

func NewDeleteTask(id int, repository repository.TaskRepository) *DeleteTask {
	return &DeleteTask{
		id:         id,
		repository: repository,
	}
}

func (u *DeleteTask) Execute() error {
	return u.repository.Delete(u.id)
}
