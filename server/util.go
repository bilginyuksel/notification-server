package main

import (
	"encoding/json"
	"net/url"
)

func hasKey(values url.Values) bool {
	return values.Get("key") == ""
}

func json2byte(v interface{}) []byte {
	bytes, _ := json.Marshal(v)
	return bytes
}
