package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mauromamani/go-clean-architecture/internal/post"
)

func MapRoutes(route *httprouter.Router, h post.Handlers) {
	route.HandlerFunc(http.MethodGet, "/api/v1/users", h.GetPosts)
	route.HandlerFunc(http.MethodPost, "/api/v1/users", h.CreatePost)
	route.HandlerFunc(http.MethodGet, "/api/v1/users/:id", h.GetPostById)
	route.HandlerFunc(http.MethodPatch, "/api/v1/users/:id", h.UpdatePost)
	route.HandlerFunc(http.MethodDelete, "/api/v1/users/:id", h.DeletePost)
}
