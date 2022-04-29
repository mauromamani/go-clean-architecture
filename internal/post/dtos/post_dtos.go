package dtos

type CreatePostDto struct {
	Title  string `json:"title" validate:"required"`
	Body   string `json:"body" validate:"required"`
	UserID int64  `json:"user_id" validate:"required"`
}

type UpdatePostDto struct {
	Title *string `json:"title" validate:"omitempty,alphanum"`
	Body  *string `json:"body" validate:"omitempty,alphanum"`
}
