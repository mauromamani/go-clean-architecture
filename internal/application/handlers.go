package application

import (
	"net/http"

	userHttp "github.com/mauromamani/go-clean-architecture/internal/user/delivery/http"
	userRepo "github.com/mauromamani/go-clean-architecture/internal/user/repository"
	userUC "github.com/mauromamani/go-clean-architecture/internal/user/usecase"

	"github.com/julienschmidt/httprouter"
)

// mapHandlers: setup all entity handlers in the application
func (s *application) mapHandlers() http.Handler {
	router := httprouter.New()

	// init repositories
	userRepository := userRepo.NewUserRepository()

	// init useCases
	userUseCase := userUC.NewUserUseCase(userRepository)

	// init http handlers
	userHandler := userHttp.NewUserHandlers(userUseCase)

	// map Routes
	userHttp.MapRoutes(router, userHandler)

	return router
}
