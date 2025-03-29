package usecase

import (
	"github.com/ksh3/go-api/src/feature/user/domain"
	"github.com/ksh3/go-api/src/feature/user/domain/entity"
)

type UserUseCase interface {
	GetSubscribeUsers() ([]entity.UserEntity, error)
}

type DefaultUserUseCase struct {
	repo           domain.UserRepository
	holidayAdapter domain.HolidayAdapter
}

func (u *DefaultUserUseCase) GetSubscribeUsers() ([]entity.UserEntity, error) {
	return nil, nil
}

func (u *DefaultUserUseCase) PreReserveAppointment() ([]string, error) {
	events, _ := u.holidayAdapter.GetEvents()
	// NOTE: Exclude holidays from the list of dates to be reserved.
	return events, nil
}

func NewDefaultUserUseCase(repo domain.UserRepository, adapter domain.HolidayAdapter) UserUseCase {
	return &DefaultUserUseCase{repo, adapter}
}
