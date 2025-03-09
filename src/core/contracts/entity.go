package contracts

type Entity[T any] interface {
	Copyble[T]
}
