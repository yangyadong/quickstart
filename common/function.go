package common

type Result struct {
	Code    int
	Message string
	Data    interface{}
}

func SendResponse(code int, message string, data interface{}) (result Result) {
	result.Code = code
	result.Message = message
	result.Data = data
	return
}
