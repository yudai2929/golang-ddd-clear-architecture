package repositories

import "golang-ddd-clear-architecture/day4/task3/domain/entities"

type UserRepository interface {
	FindAll() ([]entities.User, error)
	FindByID(id int) (entities.User, error)
	Create(user entities.User) (entities.User, error)
}
