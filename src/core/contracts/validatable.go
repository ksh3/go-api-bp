package contracts

type Validatable interface {
	Validate() error
}
