package error

// Link errors is error corresponding to link
//
// It has detail code xxx201 -> xxx299
//
// xxx is status code
//
// xxx = 400 -> Bad Request
//
// xxx = 401 -> Unauthorized
//
// xxx = 403 -> Forbidden
//
// xxx = 404 -> Not Found
//
// xxx = 429 -> Too Many Requests
//
// xxx = 500 -> Internal Server Error

// 400201 -> 400299
func UserBodyInvalid() XURLError {
	return XURLError{
		Code:       400,
		Message:    "Invalid request",
		DetailCode: 400201,
	}
}

func PasswordInvalid() XURLError {
	return XURLError{
		Code:       400,
		Message:    "Invalid password. Password must be at least 9 characters long.",
		DetailCode: 400202,
	}
}

// 401201 -> 401299
// 403201 -> 403299
// 404201 -> 404299
// 429201 -> 429299
// 500201 -> 500299
