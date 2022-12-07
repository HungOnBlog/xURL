package utils

import "encoding/json"

func InterfaceToJsonString(i interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}

func StringToInt(s string) int {
	var i int
	json.Unmarshal([]byte(s), &i)
	return i
}
