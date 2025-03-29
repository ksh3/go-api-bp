package domain

type HolidayAdapter interface {
	GetEvents() ([]string, error)
}
