package utilities

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type jsonError struct {
	Err         error  `json:"-"`
	Message     string `json:"message,omitempty"`
	jsonLogging bool   `json:"-"`
}

type ErrorProperties struct {
	Err         error  `json:"err"`
	Message     string `json:"message"`
	Fatal       bool   `json:"fatal"`
	JSONLogging bool   `json:"json_logging"`
}

type jsonErrorStringTemplate struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

func (j *jsonErrorStringTemplate) ToJSON() (jsonBytes []byte) {
	return ToJSON(j)
}

func (j *jsonErrorStringTemplate) ToJSONString() (jsonString string) {
	return ToJSONString(j)
}

func HandleError(properties ErrorProperties) *jsonError {
	if properties.Err != nil {

		if jErr, ok := properties.Err.(jsonError); ok {
			jErr.jsonLogging = properties.JSONLogging
			msgIsDifferent := (jErr.Message != properties.Message) && properties.Message != ""

			if properties.Fatal {
				if msgIsDifferent {
					log.Panic(properties.Message)
				}

				jErr.Panic("")
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
			jErr.Panic("")
		}

		log.Println(jErr.Error())

		return &jErr
	}

	return nil
}

func (e *jsonError) Panic(message string) {
	log.Println("FATAL ERROR !")
	if message != "" {
		log.Printf("Additional error message: %s\n", message)
	}
	log.Println(e.Error())
	time.Sleep(time.Second * 1)
	log.Fatal()
}

func (e jsonError) Error() string {
	if e.jsonLogging {
		x := jsonErrorStringTemplate{
			Error:   e.Err.Error(),
			Message: e.Message,
		}

		return x.ToJSONString()
	}

	messages := []string{
		e.Err.Error(),
		e.Message,
	}

	messages = MapSlice(messages, func(m string, i int) string {
		if i == 0 && len(m) > 0 {
			return fmt.Sprintf("Error: %s", m)
		} else if i == 1 && len(m) > 0 {
			return fmt.Sprintf("Message: %s", m)
		}

		return ""
	})

	messages = FilterSlice(messages, func(m string, _ int) bool {
		return len(strings.TrimSpace(m)) > 0
	})

	message := strings.TrimSpace(
		strings.Join(messages, "\n"),
	)

	return message
}

func (e jsonError) ToJSON() (jsonBytes []byte) {
	return ToJSON(e)
}

func (e jsonError) ToJSONString() (jsonString string) {
	return ToJSONString(e)
}
