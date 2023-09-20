package response

type Response struct {
	Message string
	Status  string
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

func Error(message string) Response {
	return Response{
		Message: message,
		Status:  StatusError,
	}
}
