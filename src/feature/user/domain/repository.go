package domain

import (
	"github.com/ksh3/go-api/src/feature/user/domain/entity"
)

type UserRepository interface {
	GetSubscribeUsers() ([]entity.UserEntity, error)
}
