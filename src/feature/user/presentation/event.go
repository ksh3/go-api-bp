package presentation

import (
	"reflect"
)

type UserEvent interface {
	Key() string
}

type LoadUserProfileEvent struct{}

func (e *LoadUserProfileEvent) Key() string {
	return reflect.TypeOf(*e).String()
}

type UpdateProfileEvent struct {
	Name     string `json:"name"`
	IconFile []byte `json:"iconFile"`
}

func (e *UpdateProfileEvent) Key() string {
	return reflect.TypeOf(*e).String()
}
