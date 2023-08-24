package infrastructure

import (
	"database/sql"
	"golang-ddd-clear-architecture/day4/task3/domain/entities"
	"golang-ddd-clear-architecture/day4/task3/domain/repositories"
)

type TaskRepositoryImpl struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) repositories.TaskRepository {
	return &TaskRepositoryImpl{db: db}
}

func (r TaskRepositoryImpl) FindAll() ([]entities.Task, error) {

	rows, err := r.db.Query("SELECT task_id, title, description, status, priority, created_at, update_at FROM tasks")

	if err != nil {
		return nil, err
	}

	var tasks []entities.Task

	for rows.Next() {
		var task entities.Task

		err := rows.Scan(
			&task.TaskID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.Priority,
			&task.CreatedAt,
			&task.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r TaskRepositoryImpl) FindByID(id int) (entities.Task, error) {

	task := entities.Task{}

	row := r.db.QueryRow("SELECT task_id, title, description, status, priority, created_at, update_at FROM tasks WHERE task_id = ? AND deleted_at IS NOT NULL", id)

	err := row.Scan(
		&task.TaskID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.Priority.Value,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		return entities.Task{}, err
	}

	return task, nil
}

func (r TaskRepositoryImpl) FindByUserID(id int) ([]entities.Task, error) {

	rows, err := r.db.Query("SELECT task_id, title, description, status, priority, created_at, update_at FROM tasks WHERE user_id = ? AND deleted_at IS NOT NULL", id)

	if err != nil {
		return nil, err
	}

	var tasks []entities.Task

	for rows.Next() {
		var task entities.Task

		err := rows.Scan(
			&task.TaskID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.Priority,
			&task.CreatedAt,
			&task.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r TaskRepositoryImpl) Create(task entities.Task) error {

	stmt, err := r.db.Prepare("INSERT INTO tasks (title, description, status, priority) VALUES (?, ?, ?, ?)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(task.Title, task.Description, task.Status, task.Priority)

	if err != nil {
		return err
	}

	return nil
}

func (r TaskRepositoryImpl) Update(task entities.Task) error {

	stmt, err := r.db.Prepare("UPDATE tasks SET title = ?, description = ?, status = ?, priority = ? WHERE task_id = ?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(task.Title, task.Description, task.Status, task.Priority, task.TaskID)

	if err != nil {
		return err
	}

	return nil
}

func (r TaskRepositoryImpl) Delete(id int) error {

	// 論理削除
	stmt, err := r.db.Prepare("UPDATE tasks SET deleted_at = NOW() WHERE task_id = ?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
