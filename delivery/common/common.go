package common

import "fmt"

type ResponseSuccess struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func SuccessResponse(data interface{}) ResponseSuccess {
	return ResponseSuccess{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	}
}

func ErrorResponse(status, message string) ResponseError {
	return ResponseError{
		Status:  status,
		Message: message,
	}
}

func ActivityNotFoundResponse(id uint) ResponseError {
	return ResponseError{
		Status:  "Not Found",
		Message: fmt.Sprintf("Activity with ID %d Not Found", id),
	}
}

func ToDoNotFoundResponse(id uint) ResponseError {
	return ResponseError{
		Status:  "Not Found",
		Message: fmt.Sprintf("Todo with ID %d Not Found", id),
	}
}
