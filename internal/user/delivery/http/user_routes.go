package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mauromamani/go-clean-architecture/internal/user"
)

// MapRoutes:
func MapRoutes(route *httprouter.Router, h user.Handlers) {
	route.HandlerFunc(http.MethodGet, "/api/v1/users", h.GetUsers)
	route.HandlerFunc(http.MethodPost, "/api/v1/users", h.CreateUser)
	route.HandlerFunc(http.MethodGet, "/api/v1/users/:id", h.GetUserById)
	route.HandlerFunc(http.MethodPatch, "/api/v1/users/:id", h.UpdateUser)
	route.HandlerFunc(http.MethodDelete, "/api/v1/users/:id", h.DeleteUser)
}
