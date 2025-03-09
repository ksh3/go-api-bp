package result

type Result[T any] struct {
	Value T
	Err   error
}

func Success[T any](value T) Result[T] {
	return Result[T]{Value: value, Err: nil}
}

func Failure[T any](err error) Result[T] {
	return Result[T]{Err: err}
}

func (r Result[T]) IsSuccess() bool {
	return r.Err == nil
}

func (r Result[T]) IsFailure() bool {
	return r.Err != nil
}
