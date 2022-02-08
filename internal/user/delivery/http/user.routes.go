package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mauromamani/go-clean-architecture/internal/user"
)

// MapRoutes: init each route with its http method
func MapRoutes(userGroup *gin.RouterGroup, h user.Handlers) {
	userGroup.GET("/", h.Get)
	userGroup.GET("/:id", h.GetById)
	userGroup.POST("/", h.Create)
	userGroup.PUT("/:id", h.Update)
	userGroup.DELETE("/:id", h.Delete)
}
