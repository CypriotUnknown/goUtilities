package utilities

import (
	"encoding/json"
)

type Jsonable interface {
	ToJSON() (jsonBytes []byte)
	ToJSONString() (jsonString string)
}

func FromJSON[T Jsonable](data []byte, object *T) (err error) {
	err = json.Unmarshal(data, &object)
	return err
}

func ToJSON(object any) (jsonBytes []byte) {
	jsonBytes, err := json.MarshalIndent(object, "", "    ")
	if err != nil {
		// HandleError(ErrorProperties{
		// 	Err:     err,
		// 	Message: "could not convert object to JSON",
		// 	Fatal:   true,
		// })

		return nil
	}

	return jsonBytes
}

func ToJSONString(object Jsonable) string {
	jsonBytes := object.ToJSON()
	if jsonBytes != nil {
		return string(jsonBytes)
	}

	return ""
}
