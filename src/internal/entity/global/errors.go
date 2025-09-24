package global

import (
	"errors"
	"net/http"
)

var (
	// ErrNoData данные не найдены"
	ErrNoData = errors.New("NoData")
	// ErrInternalError внутряя ошибка
	ErrInternalError = errors.New("InternalError")
	// ErrInvalidParam не валидные параметры
	ErrInvalidParam = errors.New("InvalidParam")
)

var ErrStatusCodes = map[error]int{
	ErrNoData:        http.StatusNotFound,
	ErrInternalError: http.StatusInternalServerError,
	ErrInvalidParam:  http.StatusBadRequest,
}
