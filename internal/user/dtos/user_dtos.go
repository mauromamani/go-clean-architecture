package dtos

type CreateUserDto struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,passwd"`
}

// TODO: alphanum no permite usar espacios en blanco, se puede enviar un json vacio (no modifica en nada pero hay que verlo)
type UpdateUserDto struct {
	Name     *string `json:"name" validate:"omitempty,alphanum"`
	Email    *string `json:"email" validate:"omitempty,required,email"`
	Password *string `json:"password" validate:"omitempty,required,passwd"`
}
