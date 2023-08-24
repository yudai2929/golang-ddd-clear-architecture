package requests

type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    int    `json:"priority"`
}
