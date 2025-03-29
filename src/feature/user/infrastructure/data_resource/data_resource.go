// File: data_resource.go
package data_resource

import (
	"github.com/ksh3/go-api/src/feature/user/infrastructure"
)

type DataResource interface {
	GetUsers() ([]infrastructure.UserDTO, error)
}
