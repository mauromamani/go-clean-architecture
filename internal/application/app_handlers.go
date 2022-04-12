package application

import (
	"github.com/julienschmidt/httprouter"
	userHttp "github.com/mauromamani/go-clean-architecture/internal/user/delivery/http"
	userRepo "github.com/mauromamani/go-clean-architecture/internal/user/repository"
	userUC "github.com/mauromamani/go-clean-architecture/internal/user/usecase"
)

// mapHandlers: setup all entity handlers in the application
func (app *application) mapHandlers() {
	// init repositories
	userRepository := userRepo.NewUserRepository(app.db)

	// init useCases
	userUseCase := userUC.NewUserUseCase(userRepository)

	// init http handlers
	userHandler := userHttp.NewUserHandlers(userUseCase, app.logger)

	// map Routes
	userHttp.MapRoutes(app.srv.Handler.(*httprouter.Router), userHandler)
}
