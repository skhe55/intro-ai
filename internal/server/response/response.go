package response

type Response struct {
	Message string
	Status  string
	Result  interface{}
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK(message string, result interface{}) Response {
	return Response{
		Message: message,
		Status:  StatusOK,
		Result:  &result,
	}
}

func Error(message string) Response {
	return Response{
		Message: message,
		Status:  StatusError,
		Result:  nil,
	}
}
