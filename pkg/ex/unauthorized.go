package ex

import "github.com/teytattze/learn-go-gin/pkg/status"

const unauthorizedStatus = status.BAD_REQUEST

func Unauthorized(errorCode string, message string) (int, Error) {
	err := Error{Status: unauthorizedStatus, ErrorCode: errorCode, Message: message}
	return unauthorizedStatus, err
}
