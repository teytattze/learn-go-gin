package ex

import "github.com/teytattze/learn-go-gin/pkg/status"

const notFoundStatus = status.NOT_FOUND

func NotFound(errorCode string, message string) (int, Error) {
	err := Error{Status: notFoundStatus, ErrorCode: errorCode, Message: message}
	return notFoundStatus, err
}
