package server

import (
	userHttp "github.com/mauromamani/go-clean-architecture/internal/user/delivery/http"
	userRepo "github.com/mauromamani/go-clean-architecture/internal/user/repository"
	userUC "github.com/mauromamani/go-clean-architecture/internal/user/usecase"
)

// mapHandlers: setup all entity handlers in the application
func (s *server) mapHandlers() {
	v1 := s.gin.Group("/api/v1")

	userGroup := v1.Group("/user")

	// init repositories
	userRepository := userRepo.NewUserRepository()

	// init useCases
	userUseCase := userUC.NewUserUseCase(userRepository)

	// init http handlers
	userHandler := userHttp.NewUserHandlers(userUseCase)

	// map Routes
	userHttp.MapRoutes(userGroup, userHandler)
}
