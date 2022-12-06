package error

type XURLError struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	DetailCode int    `json:"detail_code"`
	Err        error
}

func (e XURLError) Error() string {
	return "Code=" + string(e.Code) + ", Message=" + e.Message + ", DetailCode=" + string(e.DetailCode)
}
