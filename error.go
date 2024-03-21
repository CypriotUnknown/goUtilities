package utilities

type jsonError struct {
	Err     error  `json:"err"`
	Message string `json:"message"`
}

func JSONError(err error, message string) jsonError {
	return jsonError{
		Err:     err,
		Message: message,
	}
}

func (e *jsonError) Error() string {
	return ToJSONString(e)
}

func (e *jsonError) ToJSON() (jsonBytes []byte) {
	return ToJSON(e)
}

func (e *jsonError) ToJSONString() (jsonString string) {
	return ToJSONString(e)
}
