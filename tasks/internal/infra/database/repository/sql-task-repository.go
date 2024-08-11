package repository

import (
	"database/sql"
	"time"

	"github.com/gabe-frasz/gopro/tasks/internal/app/entity"
)

type SqlTaskRepository struct {
	db *sql.DB
}

func NewSqlTaskRepository(db *sql.DB) *SqlTaskRepository {
	return &SqlTaskRepository{
		db: db,
	}
}

func (r *SqlTaskRepository) Add(task *entity.Task) error {
	_, err := r.db.Exec(
		"INSERT INTO tasks (description, done, created_at) VALUES (?, ?, ?)",
		task.Description,
		task.Done,
		task.FormatCreatedAt(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *SqlTaskRepository) GetById(id int) (*entity.Task, error) {
	res := r.db.QueryRow("SELECT * FROM tasks WHERE id = ?", id)

	var task entity.Task
	var createdAt string
	err := res.Scan(&task.Id, &task.Description, &task.Done, &createdAt)
	if err != nil {
		return nil, err
	}
	task.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *SqlTaskRepository) GetAll() ([]*entity.Task, error) {
	rows, err := r.db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*entity.Task
	for rows.Next() {
		var task entity.Task
		var createdAt string
		err = rows.Scan(&task.Id, &task.Description, &task.Done, &createdAt)
		if err != nil {
			return nil, err
		}
		task.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (r *SqlTaskRepository) GetUncompleted() ([]*entity.Task, error) {
	rows, err := r.db.Query("SELECT * FROM tasks WHERE done = 0")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*entity.Task
	for rows.Next() {
		var task entity.Task
		var createdAt string
		err = rows.Scan(&task.Id, &task.Description, &task.Done, &createdAt)
		if err != nil {
			return nil, err
		}
		task.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (r *SqlTaskRepository) Update(task *entity.Task) error {
	var done int
	if task.Done {
		done = 1
	} else {
		done = 0
	}

	_, err := r.db.Exec("UPDATE tasks SET description = ?, done = ? WHERE id = ?", task.Description, done, task.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *SqlTaskRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
