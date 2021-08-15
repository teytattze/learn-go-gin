package ex

import "github.com/teytattze/learn-go-gin/pkg/status"

const badRequestStatus = status.BAD_REQUEST

func BadRequestException(errorCode string, message string) (int, Error) {
	err := Error{Status: badRequestStatus, ErrorCode: errorCode, Message: message}
	return badRequestStatus, err
}
