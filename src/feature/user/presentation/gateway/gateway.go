package gateway

import (
	"github.com/ksh3/go-api/src/core/util"

	"github.com/ksh3/go-api/src/feature/user/domain"
)

// NOTE: Use only gin handler functions in this file
type GoogleCalendarGateway struct{}

func (gateway *GoogleCalendarGateway) GetEvents() (
	[]string, error,
) {
	// NOTE: infra/service code has same logic, be DRY.
	util.GetHolidayDatesFromGoogleCalendar()
	return nil, nil
}

func NewGoogleCalendarGateway() domain.HolidayAdapter {
	return &GoogleCalendarGateway{}
}
