package contract

type Comparable[T any] interface {
	Equals(other T) bool
}
