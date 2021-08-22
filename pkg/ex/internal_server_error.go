package ex

import "github.com/teytattze/learn-go-gin/pkg/status"

const internalServerStatus = status.INTERNAL_SERVER_ERROR

func InternalServer(errorCode string, message string) (int, Error) {
	err := Error{Status: internalServerStatus, ErrorCode: errorCode, Message: message}
	return internalServerStatus, err
}
