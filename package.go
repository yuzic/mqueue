package mqueue

import (
	"encoding/json"
)

func jsonEncode(v interface{}) (string, error) {
	js, err := json.Marshal(v)
	if err != nil {
		return  "", err
	}
	return string(js), nil
}


func jsonDecode(r string, v interface{}) (error) {
	return json.Unmarshal([]byte(r), v)
}
