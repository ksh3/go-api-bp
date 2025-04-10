package contract

type Serializable[T any] interface {
	ToJSON() (string, error)
	FromJSON(jsonStr string) (T, error)
}
