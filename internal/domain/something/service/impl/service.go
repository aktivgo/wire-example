package impl

import (
	"errors"
	"wire/internal/domain/something/repository"
)

type service struct {
	repository repository.SomethingRepository
}

func NewService(
	repository repository.SomethingRepository,
) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetMessage() (string, error) {
	message := s.repository.GetMessage()

	if message == "" {
		return "", errors.New("message is empty")
	}

	return message, nil
}
