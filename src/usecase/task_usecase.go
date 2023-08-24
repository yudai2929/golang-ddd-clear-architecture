package usecase

import (
	"golang-ddd-clear-architecture/day4/task3/domain/entities"
	"golang-ddd-clear-architecture/day4/task3/domain/repositories"
	"golang-ddd-clear-architecture/day4/task3/domain/values"
	"golang-ddd-clear-architecture/day4/task3/usecase/dto"
	"golang-ddd-clear-architecture/day4/task3/usecase/params"
)

type TaskUsecase interface {
	FindAll() ([]dto.TaskDTO, error)
	FindByID(id int) (dto.TaskDTO, error)
	Create(params params.TaskParams) error
	Update(id int, params params.TaskParams) error
	Delete(id int) error
	Complete(id int) error
	InProgress(id int) error
	OnHold(id int) error
}

type TaskUsecaseImpl struct {
	taskRepository repositories.TaskRepository
}

func NewTaskUsecase(taskRepository repositories.TaskRepository) TaskUsecase {
	return &TaskUsecaseImpl{taskRepository: taskRepository}
}

func (tu TaskUsecaseImpl) FindAll() ([]dto.TaskDTO, error) {
	tasks, err := tu.taskRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var taskDTOs []dto.TaskDTO
	for _, task := range tasks {
		taskDTOs = append(taskDTOs, dto.TaskDTO{
			TaskID:      task.TaskID,
			Title:       task.Title,
			Description: task.Description,
			Priority:    task.Priority.Value,
			Status:      task.Status,
			CreatedAt:   task.CreatedAt.String(),
			UpdatedAt:   task.UpdatedAt.String(),
		})
	}

	return taskDTOs, nil
}

func (tu TaskUsecaseImpl) FindByID(id int) (dto.TaskDTO, error) {
	task, err := tu.taskRepository.FindByID(id)
	if err != nil {
		return dto.TaskDTO{}, err
	}

	return dto.TaskDTO{
		TaskID:      task.TaskID,
		Title:       task.Title,
		Description: task.Description,
		Priority:    task.Priority.Value,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt.String(),
		UpdatedAt:   task.UpdatedAt.String(),
	}, nil
}

func (tu TaskUsecaseImpl) Create(params params.TaskParams) error {
	priority, err := values.NewTaskPriority(params.Priority)
	if err != nil {
		return err
	}
	task := entities.NewTask(params.Title, params.Description, priority)
	if err := tu.taskRepository.Create(task); err != nil {
		return err
	}
	return nil
}

func (tu TaskUsecaseImpl) Update(id int, params params.TaskParams) error {
	task, err := tu.taskRepository.FindByID(id)
	if err != nil {
		return err
	}
	priority, err := values.NewTaskPriority(params.Priority)
	if err != nil {
		return err
	}

	task.Update(params.Title, params.Description, priority)
	if err := tu.taskRepository.Update(task); err != nil {
		return err
	}
	return nil

}

func (tu TaskUsecaseImpl) Delete(id int) error {
	task, err := tu.taskRepository.FindByID(id)
	if err != nil {
		return err
	}
	task.Delete()
	if err := tu.taskRepository.Delete(task.TaskID); err != nil {
		return err
	}
	return nil
}

func (tu TaskUsecaseImpl) Complete(id int) error {
	task, err := tu.taskRepository.FindByID(id)
	if err != nil {
		return err
	}
	task.Complete()
	if err := tu.taskRepository.Update(task); err != nil {
		return err
	}
	return nil
}

func (tu TaskUsecaseImpl) InProgress(id int) error {
	task, err := tu.taskRepository.FindByID(id)
	if err != nil {
		return err
	}
	task.InProgress()
	if err := tu.taskRepository.Update(task); err != nil {
		return err
	}
	return nil
}

func (tu TaskUsecaseImpl) OnHold(id int) error {
	task, err := tu.taskRepository.FindByID(id)
	if err != nil {
		return err
	}
	task.OnHold()
	if err := tu.taskRepository.Update(task); err != nil {
		return err
	}
	return nil
}
