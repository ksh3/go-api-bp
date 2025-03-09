package infrastructure

import (
	"encoding/json"
)

type userDTO struct {
	ID      *string `json:"id"`
	Name    string  `json:"name"`
	IconURL *string `json:"icon_url"`
}

type UserDTO struct {
	// NOTE: Embedding `userDTO` to reuse its fields.
	userDTO
}

type UserOpts struct {
	ID      *string
	Name    *string
	IconURL *string
}

func (u UserDTO) CopyWith(opts UserOpts) UserDTO {
	n := u
	if opts.ID != nil {
		n.userDTO.ID = opts.ID
	}
	if opts.Name != nil {
		n.userDTO.Name = *opts.Name
	}
	if opts.IconURL != nil {
		n.userDTO.IconURL = opts.IconURL
	}
	return n
}

func (u UserDTO) ToJSON() (string, error) {
	data, err := json.Marshal(u.userDTO) // `userDTO` のデータのみ JSON 化
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (u *UserDTO) FromJSON(jsonStr string) (UserDTO, error) {
	var newUser userDTO
	err := json.Unmarshal([]byte(jsonStr), &newUser)
	if err != nil {
		return UserDTO{}, err
	}
	return UserDTO{newUser}, nil
}

func NewUserDTO(opts UserOpts) UserDTO {
	return UserDTO{userDTO{
		ID:      opts.ID,
		Name:    *opts.Name,
		IconURL: opts.IconURL,
	}}
}
