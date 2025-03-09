package usecase

import (
	"github.com/ksh3/go-api/src/feature/user/domain"
	"github.com/ksh3/go-api/src/feature/user/domain/entity"
)

type UserUseCase interface {
	GetSubscribeUsers() ([]entity.UserEntity, error)
}

type DefaultUserUseCase struct {
	repo domain.UserRepository
}

func (u *DefaultUserUseCase) GetSubscribeUsers() ([]entity.UserEntity, error) {
	return nil, nil
}

func NewDefaultUserUseCase(repo domain.UserRepository) UserUseCase {
	return &DefaultUserUseCase{repo}
}
