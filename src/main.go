package main

import "golang-ddd-clear-architecture/day4/task3/adapter"

type TaskResponse struct {
	TaskID      int    `json:"task_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    int    `json:"priority"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func main() {
	adapter.InitRouter()
}
