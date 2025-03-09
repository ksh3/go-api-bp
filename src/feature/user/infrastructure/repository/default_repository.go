// File: default_repository.go
package repository

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ksh3/go-api/src/feature/user/domain"
	"github.com/ksh3/go-api/src/feature/user/domain/entity"
)

type DefaultUserRepository struct {
	db *mongo.Database
}

func (repo *DefaultUserRepository) GetSubscribeUsers() (
	[]entity.UserEntity, error,
) {
	panic("not implemented") // TODO: Implement
}

func NewDefaultUserRepository(db *mongo.Database) domain.UserRepository {
	return &DefaultUserRepository{db: db}
}
