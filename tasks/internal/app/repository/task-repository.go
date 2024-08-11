package repository

import "github.com/gabe-frasz/gopro/tasks/internal/app/entity"

type TaskRepository interface {
	Add(task *entity.Task) error
	GetById(id int) (*entity.Task, error)
	GetAll() ([]*entity.Task, error)
	GetUncompleted() ([]*entity.Task, error)
	Update(task *entity.Task) error
	Delete(id int) error
}
