package helpers

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func CreateError(status int, message string) *Error {
	return &Error{
		Status:  status,
		Message: message,
	}
}
