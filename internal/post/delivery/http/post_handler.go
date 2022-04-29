package http

import (
	"fmt"
	"net/http"

	"github.com/mauromamani/go-clean-architecture/internal/post"
	"github.com/mauromamani/go-clean-architecture/internal/post/dtos"
	httpErrors "github.com/mauromamani/go-clean-architecture/pkg/errors"
	"github.com/mauromamani/go-clean-architecture/pkg/logger"
	"github.com/mauromamani/go-clean-architecture/pkg/utils"
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
func (h *postHandlers) GetPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	posts, err := h.useCase.GetPosts(ctx)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"posts": posts}, nil)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
	}
}

// GetPostById:
func (h *postHandlers) GetPostById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ReadIDParam(r)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
		return
	}

	p, err := h.useCase.GetPostById(ctx, id)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"post": p}, nil)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
	}
}

// CreatePost:
func (h *postHandlers) CreatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	createPostDto := &dtos.CreatePostDto{}

	err := utils.ReadJSON(w, r, createPostDto)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
		return
	}

	post, err := h.useCase.CreatePost(ctx, createPostDto)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/api/v1/posts/%d", post.ID))

	err = utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"post": post}, headers)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
	}
}

// UpdatePost:
func (h *postHandlers) UpdatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ReadIDParam(r)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
		return
	}

	updatePostDto := &dtos.UpdatePostDto{}

	err = utils.ReadJSON(w, r, updatePostDto)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
		return
	}

	updatedPost, err := h.useCase.UpdatePost(ctx, id, updatePostDto)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"post": updatedPost}, nil)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
	}
}

// DeletePost:
func (h *postHandlers) DeletePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ReadIDParam(r)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
		return
	}

	err = h.useCase.DeletePost(ctx, id)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"post": "post deleted!"}, nil)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
	}
}
