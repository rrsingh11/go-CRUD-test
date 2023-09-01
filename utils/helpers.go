package utils

import (
	"encoding/json"
)


func EncodeMessage(s string) []byte {
	resp := make(map[string]string)
	resp["message"] = s
	res, _ := json.Marshal(resp)
	return res
}
