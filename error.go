package utilities

import "log"

type jsonError struct {
	Err     error  `json:"err"`
	Message string `json:"message"`
}

func HandleError(err error, message string, fatal bool) {
	if err != nil {
		jErr := jsonError{
			Err:     err,
			Message: message,
		}

		if fatal {
			jErr.Panic()
		}

		log.Println(jErr.Error())
	}
}

func (e *jsonError) Panic() {
	log.Fatalln(e.Error())
}

func (e jsonError) Error() string {
	return ToJSONString(e)
}

func (e jsonError) ToJSON() (jsonBytes []byte) {
	return ToJSON(e)
}

func (e jsonError) ToJSONString() (jsonString string) {
	return ToJSONString(e)
}
