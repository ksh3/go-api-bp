package service

// NOTE: e.g.infra mattermost, algolia, elasticsearch, RAG service, etc.
import (
	"github.com/ksh3/go-api/src/core/util"
	"github.com/ksh3/go-api/src/feature/user/domain"
)

// NOTE: Use domain layer in this file.
type GoogleCalendarService struct{}

func (service *GoogleCalendarService) GetEvents() (
	[]string, error,
) {
	events := util.GetHolidayDatesFromGoogleCalendar()
	// NOTE: Exclude holidays from the list of dates to be reserved.
	return events, nil
}

func NewGoogleCalendarService() domain.HolidayAdapter {
	return &GoogleCalendarService{}
}
