package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mauromamani/go-clean-architecture/internal/user"
)

// MapRoutes: init each route with its http method
func MapRoutes(route *httprouter.Router, h user.Handlers) {
	route.HandlerFunc(http.MethodGet, "/api/v1/users", h.GetUser)
	route.HandlerFunc(http.MethodGet, "/api/v1/users/:id", h.GetUserById)
	route.HandlerFunc(http.MethodPost, "/api/v1/users", h.CreateUser)
	route.HandlerFunc(http.MethodPatch, "/api/v1/users", h.UpdateUser)
	route.HandlerFunc(http.MethodDelete, "/api/v1/users", h.DeleteUser)
}
