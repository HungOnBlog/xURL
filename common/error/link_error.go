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
func OriginalLinkRequired() XURLError {
	return XURLError{
		Code:       400,
		Message:    "Original link is required",
		DetailCode: 400001,
	}
}

func OriginalLinkInvalid() XURLError {
	return XURLError{
		Code:       400,
		Message:    "Original link is invalid",
		DetailCode: 400002,
	}
}

func LinkRequestInvalid() XURLError {
	return XURLError{
		Code:       400,
		Message:    "Invalid request",
		DetailCode: 400001,
	}
}

// 401001 -> 401099
// 403001 -> 403099
// 404001 -> 404099
func LinkNotFound() XURLError {
	return XURLError{
		Code:       404,
		Message:    "Link not found",
		DetailCode: 404001,
	}
}

// 429001 -> 429099
// 500001 -> 500099
