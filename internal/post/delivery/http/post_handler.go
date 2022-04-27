package http

import (
	"net/http"

	"github.com/mauromamani/go-clean-architecture/internal/post"
	"github.com/mauromamani/go-clean-architecture/pkg/logger"
)

type postHandlers struct {
	useCase post.UseCase
	logger  logger.Logger
}

func NewPostHandlers(useCase post.UseCase, logger logger.Logger) post.Handlers {
	return &postHandlers{
		useCase: useCase,
		logger:  logger,
	}
}

// GetPosts:
func (h *postHandlers) GetPosts(w http.ResponseWriter, r *http.Request) {}

// GetPostById:
func (h *postHandlers) GetPostById(w http.ResponseWriter, r *http.Request) {}

// CreatePost:
func (h *postHandlers) CreatePost(w http.ResponseWriter, r *http.Request) {}

// UpdatePost:
func (h *postHandlers) UpdatePost(w http.ResponseWriter, r *http.Request) {}

// DeletePost:
func (h *postHandlers) DeletePost(w http.ResponseWriter, r *http.Request) {}
