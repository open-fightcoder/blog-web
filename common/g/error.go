package g

import "net/http"

type Error struct {
	Code     int
	HttpCode int
	Msg      string
}

func _build(httpCode int, code int, defval string, custom ...string) Error {
	msg := defval
	if len(custom) > 0 {
		msg = custom[0]
	}

	return Error{
		Code:     code,
		HttpCode: httpCode,
		Msg:      msg,
	}
}

func DBError(msg ...string) Error {
	return _build(http.StatusInternalServerError, 10001, "DB Error", msg...)
}

func ParamError(msg ...string) Error {
	return _build(http.StatusBadRequest, 10002, "Param Error", msg...)
}

func ServerError(msg ...string) Error {
	return _build(http.StatusInternalServerError, 10003, "Server Error", msg...)
}

func PrivError(msg ...string) Error {
	return _build(http.StatusForbidden, 10004, "Forbidden", msg...)
}

func NotFoundError(msg ...string) Error {
	return _build(http.StatusNotFound, 10005, "Not Found", msg...)
}

func RequestError(msg ...string) Error {
	return _build(http.StatusBadRequest, 10006, "Bad Request", msg...)
}
