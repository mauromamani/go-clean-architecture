package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mauromamani/go-clean-architecture/internal/user"
)

// MapRoutes: init each route with its http method
func MapRoutes(userGroup *gin.RouterGroup, h user.Handlers) {
	userGroup.GET("/", h.GetUser)
	userGroup.GET("/:id", h.GetUserById)
	userGroup.POST("/", h.CreateUser)
	userGroup.PATCH("/:id", h.UpdateUser)
	userGroup.DELETE("/:id", h.DeleteUser)
}
