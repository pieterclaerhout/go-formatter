package ydformatter

import (
	"bytes"
	"encoding/json"
)

// JSONBytes pretty prints a slice of JSON bytes
func JSONBytes(data []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", "    ")
	if err == nil {
		return string(out.Bytes())
	}
	return string(data)
}

// JSONString pretty prints a JSON string
func JSONString(data string) string {
	return JSONBytes([]byte(data))
}
