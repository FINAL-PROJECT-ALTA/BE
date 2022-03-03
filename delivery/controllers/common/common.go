package common

import "net/http"

type Response struct {
	Code    interface{} `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(code interface{}, msg interface{}, data interface{}) Response {
	if code == nil {
		code = http.StatusOK
	}
	if msg == nil {
		msg = "success"
	}
	if data == nil {
		data = nil
	}
	return Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

func Update(code interface{}, msg interface{}) Response {
	if code == nil {
		code = http.StatusOK
	}
	if msg == nil {
		msg = "success"
	}
	return Response{
		Code:    code,
		Message: msg,
	}
}

func InternalServerError(code interface{}, msg interface{}, data interface{}) Response {
	if code == nil {
		code = http.StatusInternalServerError
	}
	if msg == nil {
		msg = "error in server"
	}
	if data == nil {
		data = nil
	}
	return Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

func BadRequest(code interface{}, msg interface{}, data interface{}) Response {
	if code == nil {
		code = http.StatusBadRequest
	}
	if msg == nil {
		msg = "error in request"
	}
	if data == nil {
		data = nil
	}
	return Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}
