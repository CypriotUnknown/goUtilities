package utilities

import "log"

type jsonError struct {
	Err         error  `json:"err"`
	Message     string `json:"message,omitempty"`
	jsonLogging bool   `json:"-"`
}

type ErrorProperties struct {
	Err         error  `json:"err"`
	Message     string `json:"message"`
	Fatal       bool   `json:"fatal"`
	JSONLogging bool   `json:"json_logging"`
}

func HandleError(properties ErrorProperties) *jsonError {
	if properties.Err != nil {

		if jErr, ok := properties.Err.(jsonError); ok {
			jErr.jsonLogging = properties.JSONLogging
			msgIsDifferent := (jErr.Message != properties.Message) && properties.Message != ""

			if properties.Fatal {
				if msgIsDifferent {
					log.Fatalf("Additional error message: %s", properties.Message)
				}

				log.Fatal()
			}

			if msgIsDifferent {
				log.Printf("Additional error message: %s", properties.Message)
			}

			return &jErr
		}

		jErr := jsonError{
			Err:         properties.Err,
			Message:     properties.Message,
			jsonLogging: properties.JSONLogging,
		}

		if properties.Fatal {
			jErr.Panic()
		}

		log.Println(jErr.Error())

		return &jErr
	}

	return nil
}

func (e *jsonError) Panic() {
	log.Fatalln(e.Error())
}

func (e jsonError) Error() string {
	if e.jsonLogging {
		return ToJSONString(e)
	}

	return e.Message
}

func (e jsonError) ToJSON() (jsonBytes []byte) {
	return ToJSON(e)
}

func (e jsonError) ToJSONString() (jsonString string) {
	return ToJSONString(e)
}
