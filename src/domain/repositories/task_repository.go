package repositories

import "golang-ddd-clear-architecture/day4/task3/domain/entities"

type TaskRepository interface {
	FindAll() ([]entities.Task, error)
	FindByID(id int) (entities.Task, error)
	FindByUserID(userID int) ([]entities.Task, error)
	Create(task entities.Task) error
	Update(task entities.Task) error
	Delete(id int) error
}
