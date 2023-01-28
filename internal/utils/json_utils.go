package utils

import "encoding/json"

func InterfaceToString(obj interface{}) string {
	bytes, _ := json.MarshalIndent(obj, "", "\t")
	return string(bytes)
}
