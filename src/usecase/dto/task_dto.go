package dto

type TaskDTO struct {
	TaskID      int    `json:"task_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
