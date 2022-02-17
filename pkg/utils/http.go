package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/mauromamani/go-clean-architecture/pkg/validator"
)

// ReadRequest: Read request and validate
func ReadRequest(c *gin.Context, req interface{}) error {
	if err := c.BindJSON(req); err != nil {
		return err
	}

	return validator.ValidateStruct(c.Request.Context(), req)
}
