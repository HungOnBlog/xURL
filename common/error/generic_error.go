package error

// Link errors is error corresponding to link
//
// It has detail code xxx100 -> xxx199
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

// 400101 -> 400199
func EmailInvalid() XURLError {
	return XURLError{
		Code:       400,
		Message:    "Email is invalid",
		DetailCode: 400101,
	}
}

// 401101 -> 401199
func ApikeyInvalid() XURLError {
	return XURLError{
		Code:       401,
		Message:    "Apikey is invalid or empty. Please check your apikey or create a new one",
		DetailCode: 401101,
	}
}

func UnauthorizedUserNotFound() XURLError {
	return XURLError{
		Code:       401,
		Message:    "Unauthorized user not found",
		DetailCode: 401102,
	}
}

// 403101 -> 403199
// 404101 -> 404199
// 429101 -> 429199
// 500101 -> 500199
func InternalServerError() error {
	return XURLError{
		Code:       500,
		Message:    "Internal server error",
		DetailCode: 500001,
	}
}
