package user

import "github.com/gin-gonic/gin"

type Handlers interface {
	GetUser(c *gin.Context)
	GetUserById(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
