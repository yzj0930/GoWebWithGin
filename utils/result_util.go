package util

import "github.com/yzj0930/GoWebWithGin/dto/response"

const SUCCESS_CODE = 0
const DEFAULT_ERROR_CODE = -1
const SUCCESS_MSG = "Success"
const ERROR_MSG = "Error"

func ReturnResult(status int, message string, data interface{}) response.ResponseDto {
	return response.ResponseDto{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func ReturnSuccess(data interface{}) response.ResponseDto {
	return ReturnResult(SUCCESS_CODE, SUCCESS_MSG, data)
}

func ReturnError(message string) response.ResponseDto {
	return ReturnResult(DEFAULT_ERROR_CODE, ERROR_MSG, nil)
}

func ReturnErrorE(err error) response.ResponseDto {
	return ReturnResult(DEFAULT_ERROR_CODE, err.Error(), nil)
}

func ReturnErrorWithStatus(status int, message string) response.ResponseDto {
	return ReturnResult(status, message, nil)
}
