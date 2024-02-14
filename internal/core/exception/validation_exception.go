package exception

type ValidationException struct {
	Message string
}

func NewValidationException(message string) *ValidationException {
	return &ValidationException{
		Message: message,
	}
}

func (e *ValidationException) Error() string {
	return e.Message
}
