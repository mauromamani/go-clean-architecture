package user

type UseCase interface {
	Get()
	GetById()
	Create()
	Update()
	Delete()
}
