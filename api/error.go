package api

import (
	"encoding/json"
	"fmt"
)

type e struct {
	Code int    `json:"error_code"`
	Text string `json:"error_text"`
}

// Error wraps an error in a struct to encode in JSON
func Error(err error) []byte {
	if err != nil {
		t := &e{Code: 1, Text: fmt.Sprintf("%v", err)}
		j, _ := json.Marshal(t)
		return j
	}
	// FIXME: Return nothing instead?
	return []byte("{}")
}
