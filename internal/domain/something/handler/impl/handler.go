package impl

import (
	"log"
	"wire/internal/domain/something/service"
)

type handler struct {
	service service.Service
}

func NewHandler(
	service service.Service,
) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Handle() error {
	message, err := h.service.GetMessage()
	if err != nil {
		return err
	}

	log.Println("repository message:", message)

	return nil
}
