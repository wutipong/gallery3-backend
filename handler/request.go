package handler

import (
	"encoding/json"
	"io"
)

func ParseInput(r io.Reader, value interface{}) error {
	reqBody, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	if len(reqBody) != 0 {
		return json.Unmarshal(reqBody, value)
	}

	return nil
}
