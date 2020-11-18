package core

import "encoding/json"

func FormatOutput(code int, msg string, data interface{}) string {
	output, _ := json.Marshal(map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	})
	return string(output)
}

func ConvertBoolToInt(temp bool) int {
	if temp {
		return 1
	}
	return 0
}
