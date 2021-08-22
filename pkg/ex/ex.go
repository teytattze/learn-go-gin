package ex

type Error struct {
	Status    int    `json:"status"`
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
}

func Http(status int, errorCode string, message string) (int, Error) {
	err := Error{Status: status, ErrorCode: errorCode, Message: message}
	return status, err
}
