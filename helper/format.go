package helper

import "encoding/json"

func FormatStruct(v interface{}) string {
	vjson, _ := json.MarshalIndent(v, "", "  ")
	return string(vjson)
}