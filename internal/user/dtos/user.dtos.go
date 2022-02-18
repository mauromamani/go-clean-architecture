package dtos

type CreateUserDto struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,passwd"`
}

type UpdateUserDto struct {
	Name     *string `json:"name" validate:"omitempty,required"`
	Email    *string `json:"email" validate:"omitempty,required,email"`
	Password *string `json:"password" validate:"omitempty,required,passwd"`
}
