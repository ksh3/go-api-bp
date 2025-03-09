package entity

type UserEntity struct {
	ID      string  `json:"id" validate:"required"`
	Name    string  `json:"name" validate:"required"`
	IconURL *string `json:"icon_url" validate:"omitempty"`
	HasIcon bool    `json:"has_icon" validate:"required"`
}
