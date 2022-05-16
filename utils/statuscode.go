package utils

type Status_code struct {
	Code    int
	Message string
}

func NotFound() Status_code {
	return Status_code{Code: 404, Message: "Not Found! Please try again."}
}

func Ok() Status_code {
	return Status_code{Code: 200, Message: "Ok"}
}

func BadRequest() Status_code {
	return Status_code{Code: 500, Message: "Failed to process request.Please try again!"}
}
