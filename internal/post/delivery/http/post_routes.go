package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mauromamani/go-clean-architecture/internal/post"
)

func MapRoutes(route *httprouter.Router, h post.Handlers) {
	route.HandlerFunc(http.MethodGet, "/api/v1/posts", h.GetPosts)
	route.HandlerFunc(http.MethodPost, "/api/v1/posts", h.CreatePost)
	route.HandlerFunc(http.MethodGet, "/api/v1/posts/:id", h.GetPostById)
	route.HandlerFunc(http.MethodPatch, "/api/v1/posts/:id", h.UpdatePost)
	route.HandlerFunc(http.MethodDelete, "/api/v1/posts/:id", h.DeletePost)
}
