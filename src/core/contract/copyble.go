package contract

type Copyble[T any] interface {
	CopyWith(opts T) T
}
