package contracts

type DTO[T any] interface {
	Copyble[T]
	FromJson(json string) T
	ToJson() string
}
