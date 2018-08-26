package base

type HttpFailResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func Fail(errCode int, errMsg string) *HttpFailResponse {
	return &HttpFailResponse{Code: errCode, Msg: errMsg}
}
