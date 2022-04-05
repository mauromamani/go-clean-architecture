package utils

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
