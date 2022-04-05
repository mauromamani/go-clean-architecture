package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mauromamani/go-clean-architecture/internal/user"
)

// MapRoutes: init each route with its http method
func MapRoutes(route *httprouter.Router, h user.Handlers) {
	route.HandlerFunc(http.MethodGet, "/api/v1/users", h.GetUser)
}
