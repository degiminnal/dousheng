package service

type ResponseCommon struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}
