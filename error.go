package utilities

type JsonError struct {
	Err     error  `json:"err"`
	Message string `json:"message"`
}

func JSONError(err error, message string) JsonError {
	return JsonError{
		Err:     err,
		Message: message,
	}
}

func (e JsonError) Error() string {
	return ToJSONString(e)
}

func (e JsonError) ToJSON() (jsonBytes []byte) {
	return ToJSON(e)
}

func (e JsonError) ToJSONString() (jsonString string) {
	return ToJSONString(e)
}
