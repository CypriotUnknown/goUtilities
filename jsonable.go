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
	HandleError(err, "could not convert object to JSON", true)

	return jsonBytes
}

func ToJSONString(object Jsonable) string {
	return string(object.ToJSON())
}
