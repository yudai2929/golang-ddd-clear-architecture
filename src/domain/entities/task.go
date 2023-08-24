package entities

import (
	"errors"
	"golang-ddd-clear-architecture/day4/task3/domain/fields"
	"golang-ddd-clear-architecture/day4/task3/domain/values"
	"time"
)

type Task struct {
	TaskID      int
	Title       string
	Description string
	Priority    values.TaskPriority
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeleteAt    *time.Time
}

func NewTask(title string, description string, priority values.TaskPriority) Task {
	return Task{
		Title:       title,
		Description: description,
		Priority:    priority,
		Status:      fields.TaskStatus.InProgress,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeleteAt:    nil,
	}
}

func (t *Task) Complete() error {
	if t.Status == fields.TaskStatus.Completed {
		return errors.New("task already completed")
	}
	t.Status = fields.TaskStatus.Completed
	t.UpdatedAt = time.Now()

	return nil
}

func (t *Task) InProgress() error {
	t.Status = fields.TaskStatus.InProgress
	t.UpdatedAt = time.Now()

	return nil
}

func (t *Task) OnHold() error {
	if t.Status == fields.TaskStatus.Completed {
		return errors.New("cannot set on hold for completed task")
	}
	t.Status = fields.TaskStatus.OnHold
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) Delete() {
	t.DeleteAt = new(time.Time)
	*t.DeleteAt = time.Now()
}

func (t *Task) Update(title string, description string, priority values.TaskPriority) {
	t.Title = title
	t.Description = description
	t.Priority = priority
	t.UpdatedAt = time.Now()
}

func (t *Task) IsDeleted() bool {
	return t.DeleteAt != nil
}

func (t *Task) IsCompleted() bool {
	return t.Status == fields.TaskStatus.Completed
}

func (t *Task) IsInProgress() bool {
	return t.Status == fields.TaskStatus.InProgress
}

func (t *Task) IsHighPriority() bool {
	return t.Priority.Value == 5 || t.Priority.Value == 4
}
