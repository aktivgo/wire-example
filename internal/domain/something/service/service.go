package service

type Service interface {
	GetMessage() (string, error)
}
