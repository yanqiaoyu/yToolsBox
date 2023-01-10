package dto

type FailResponseMeta struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type SuccessResponseMeta struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
