package usecases

import (
	"github.com/gabe-frasz/gopro/tasks/internal/app/repository"
)

type CompleteTask struct {
	id         int
	repository repository.TaskRepository
}

func NewCompleteTask(id int, repository repository.TaskRepository) *CompleteTask {
	return &CompleteTask{
		id:         id,
		repository: repository,
	}
}

func (c *CompleteTask) Execute() error {
	task, err := c.repository.GetById(c.id)
	if err != nil {
		return err
	}

	task.Complete()
	return c.repository.Update(task)
}
