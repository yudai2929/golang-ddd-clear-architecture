package values

import "fmt"

type TaskPriority struct {
	Value int
}

func NewTaskPriority(value int) (TaskPriority, error) {
	if value < 1 || value > 5 {
		return TaskPriority{}, fmt.Errorf("invalid priority value: %d", value)
	}
	return TaskPriority{Value: value}, nil
}
