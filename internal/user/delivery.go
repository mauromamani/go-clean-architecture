package user

import "github.com/gin-gonic/gin"

type Handlers interface {
	Get(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
