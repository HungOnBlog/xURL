package error

import "encoding/json"

type XURLError struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	DetailCode int    `json:"detail_code"`
	Err        error
}

func (e XURLError) Error() string {
	str, _ := json.Marshal(e)
	return string(str)
}
