package dtos

type CreatePostDto struct {
	Title  string `json:"title" validate:"required,trim"`
	Body   string `json:"body" validate:"required,trim"`
	UserID int64  `json:"user_id" validate:"required,trim"`
}

type UpdatePostDto struct {
	Title *string `json:"title" validate:"omitempty,trim"`
	Body  *string `json:"body" validate:"omitempty,trim"`
}
