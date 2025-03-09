// File: remote_resource.go
package resource

import (
	"github.com/ksh3/go-api/src/feature/user/infrastructure"
	"go.mongodb.org/mongo-driver/mongo"
)

type RemoteUserResource struct {
	collection *mongo.Collection
}

func (r *RemoteUserResource) GetUsers() ([]infrastructure.UserDTO, error) {
	return nil, nil
}

func NewRemoteUserResource(collection *mongo.Collection) DataResource {
	return &RemoteUserResource{collection}
}
