package main

import (
	"github.com/google/wire"
	"log"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Handle() {
	log.Println(h.service.repository.Message)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

type Repository struct {
	Message string
}

func NewRepository(message string) *Repository {
	return &Repository{
		Message: message,
	}
}

func InitializeHandler(message string) Handler {
	panic(wire.Build(NewRepository(message), NewService, NewHandler))
}

func main() {
	handler := InitializeHandler("Hello world")
	handler.Handle()
}
