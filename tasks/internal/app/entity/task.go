package entity

import "time"

type Task struct {
	Id          int
	Description string
	Done        bool
	CreatedAt   time.Time
}

func NewTask(id int, description string, done bool) *Task {
	return &Task{
		Id:          id,
		Description: description,
		Done:        done,
		CreatedAt:   time.Now(),
	}
}

func (t *Task) Complete() {
	t.Done = true
}

func (t *Task) FormatCreatedAt() string {
	return t.CreatedAt.Format(time.RFC3339)
}
