package ex

import "github.com/teytattze/learn-go-gin/pkg/status"

const internalServerStatus = status.BAD_REQUEST

func InternalServerException(errorCode string, message string) (int, Error) {
	err := Error{Status: internalServerStatus, ErrorCode: errorCode, Message: message}
	return badRequestStatus, err
}
