package presentation

import (
	"reflect"
)

type UserState interface {
	Key() string
}

type LoadedUserProfileState struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	IconURL string `json:"iconURL"`
}

func (e *LoadedUserProfileState) Key() string {
	return reflect.TypeOf(*e).String()
}

type LoadingUserProfileState struct{}

func (e *LoadingUserProfileState) Key() string {
	return reflect.TypeOf(*e).String()
}
