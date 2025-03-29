// File: local_resource.go
package data_resource

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/ksh3/go-api/src/feature/user/infrastructure"
)

type LocalUserResource struct {
	db *badger.DB
}

func (r *LocalUserResource) GetUsers() ([]infrastructure.UserDTO, error) {
	return nil, nil
}

func NewLocalUserResource(db *badger.DB) DataResource {
	return &LocalUserResource{db}
}
