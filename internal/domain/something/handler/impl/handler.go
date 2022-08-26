package impl

import (
	"log"
	"wire/internal/domain/something/service"
)

type handler struct {
	service service.SomethingService
}

func NewHandler(
	service service.SomethingService,
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
