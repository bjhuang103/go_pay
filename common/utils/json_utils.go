package utils

import "encoding/json"

func JsonMarshal(data interface{}) string {
	b, _ := json.Marshal(data)
	return string(b)
}
