package user

type Repository interface {
	Get()
	GetById()
	Create()
	Update()
	Delete()
}
