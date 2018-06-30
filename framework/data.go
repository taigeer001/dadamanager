package xf

import (
	"encoding/json"
	"fmt"
)

func ToJson(data map[string]interface{}) []byte {
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return []byte("")
	}
	return b
}