package common

import (
	"testing"

	"github.com/go-playground/assert"
)

type MockResponse struct {
	Code    interface{} `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}
type MockResponseMenu struct {
	Code          interface{} `json:"code"`
	Message       interface{} `json:"message"`
	TotalResult   interface{} `json:"totalresult"`
	LimitCalories interface{} `json:"limitcalories"`
	Data          interface{} `json:"data"`
}

func (MockResponse) Success(code interface{}, msg interface{}, data interface{}) Response {
	return Response{}
}
func (MockResponse) Update(code interface{}, msg interface{}, data interface{}) Response {
	return Response{}
}
func (MockResponse) InternalServerError(code interface{}, msg interface{}, data interface{}) Response {
	return Response{}
}
func (MockResponse) BadRequest(code interface{}, msg interface{}, data interface{}) Response {
	return Response{}
}
func (MockResponseMenu) SuccessMenu(code interface{}, msg interface{}, data interface{}) ResponseMenu {
	return ResponseMenu{}
}

func TestSuccess(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, Response{Code: 200, Message: "success", Data: nil}, Success(nil, nil, nil))
	})
}
func TestUpdate(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, Response{Code: 200, Message: "success"}, Update(nil, nil))
	})
}
func TestInternalServerError(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, Response{Code: 500, Message: "error in server", Data: nil}, InternalServerError(nil, nil, nil))
	})
}
func TestBadRequest(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, Response{Code: 400, Message: "error in request", Data: nil}, BadRequest(nil, nil, nil))
	})
}
func TestSuccessMenu(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, ResponseMenu{Code: 200, Message: "success", Data: nil}, SuccessMenu(nil, nil, nil, nil, nil))
	})
}
