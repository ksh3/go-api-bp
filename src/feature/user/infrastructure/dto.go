package infrastructure

// emebedding
type userDTO struct {
	ID      *string
	Name    string
	IconURL *string
}

type UserDTO struct {
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
