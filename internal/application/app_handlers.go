package application

import (
	"github.com/julienschmidt/httprouter"
	postHttp "github.com/mauromamani/go-clean-architecture/internal/post/delivery/http"
	postRepo "github.com/mauromamani/go-clean-architecture/internal/post/repository"
	postUC "github.com/mauromamani/go-clean-architecture/internal/post/usecase"
	userHttp "github.com/mauromamani/go-clean-architecture/internal/user/delivery/http"
	userRepo "github.com/mauromamani/go-clean-architecture/internal/user/repository"
	userUC "github.com/mauromamani/go-clean-architecture/internal/user/usecase"
)

// mapHandlers: setup all entity handlers in the application
func (app *application) mapHandlers() {
	// init repositories
	userRepository := userRepo.NewUserRepository(app.db)
	postRepository := postRepo.NewPostRepository(app.db)

	// init useCases
	userUseCase := userUC.NewUserUseCase(userRepository)
	postUseCase := postUC.NewPostUseCase(postRepository)

	// init http handlers
	userHandler := userHttp.NewUserHandlers(userUseCase, app.logger)
	postHandler := postHttp.NewPostHandlers(postUseCase, app.logger)

	// map routes
	userHttp.MapRoutes(app.srv.Handler.(*httprouter.Router), userHandler)
	postHttp.MapRoutes(app.srv.Handler.(*httprouter.Router), postHandler)
}
