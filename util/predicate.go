package util

func Negate[T any](predicate func(T) bool) func(T) bool {
	return func(value T) bool {
		return !predicate(value)
	}
}
