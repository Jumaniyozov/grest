package user

import (
	"context"
	"github.com/jumaniyozov/grest/pkg/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (u User, err error) {
	// TODO need to finish create
	return User{}, nil
}
