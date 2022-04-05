package utils

import (
	"encoding/json"
	"net/http"
)

type envelope map[string]interface{}

// ReadRequest: Read request and validate
// func ReadRequest(c *gin.Context, req interface{}) error {
// 	maxBytes := 1_048_576

// 	var w http.ResponseWriter = c.Writer
// 	c.Request.Body = http.MaxBytesReader(w, c.Request.Body, int64(maxBytes))
// 	if err := c.BindJSON(req); err != nil {
// 		return err
// 	}

// 	return validator.ValidateStruct(c.Request.Context(), req)
// }

func WriteJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')

	// Include any headers
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
