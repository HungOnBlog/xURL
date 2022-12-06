package utils

import "encoding/json"

func InterfaceToJsonString(i interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}
