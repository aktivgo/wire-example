package impl

import (
	"errors"
	"wire/internal/domain/something/repository"
)

type service struct {
	repository repository.Repository
}

func NewService(
	repository repository.Repository,
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
