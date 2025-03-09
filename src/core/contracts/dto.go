package contracts

type DTO[T any] interface {
	Copyble[T]
	Serializable[T]
}
