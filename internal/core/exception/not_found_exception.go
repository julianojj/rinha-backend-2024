package exception

type NotFoundException struct {
	Message string
}

func NewNotFoundException(message string) *NotFoundException {
	return &NotFoundException{
		Message: message,
	}
}

func (e *NotFoundException) Error() string {
	return e.Message
}
