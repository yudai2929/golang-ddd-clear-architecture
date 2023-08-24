package messages

type successMessage struct {
	Message string `json:"message"`
}

func SuccessMessage(message string) successMessage {
	return successMessage{
		Message: message,
	}
}
