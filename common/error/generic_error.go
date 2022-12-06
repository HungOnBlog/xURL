package error

// Link errors is error corresponding to link
//
// It has detail code xxx001 -> xxx099
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

// 400001 -> 400099
// 401001 -> 401099
// 403001 -> 403099
// 404001 -> 404099
// 429001 -> 429099
// 500001 -> 500099
func InternalServerError() error {
	return XURLError{
		Code:       500,
		Message:    "Internal server error",
		DetailCode: 500001,
	}
}
