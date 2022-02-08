package server

import (
	userHttp "github.com/mauromamani/go-clean-architecture/internal/user/delivery/http"
)

// mapHandlers: it setup all entity handlers in the application
func (s *server) mapHandlers() {
	v1 := s.gin.Group("/api/v1")

	userGroup := v1.Group("/user")

	// Init http handlers
	userHandler := userHttp.NewUserHandlers()

	// Map Routes
	userHttp.MapRoutes(userGroup, userHandler)
}
