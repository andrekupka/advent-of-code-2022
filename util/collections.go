package util

const NotFound = -1

func FirstIndexMatching[T any](slice []T, predicate func(T) bool) int {
	for index, element := range slice {
		if predicate(element) {
			return index
		}
	}
	return NotFound
}

func LastIndexMatching[T any](slice []T, predicate func(T) bool) int {
	for index := len(slice) - 1; index >= 0; index-- {
		if predicate(slice[index]) {
			return index
		}
	}
	return NotFound
}

func DropWhile[T any](slice []T, predicate func(T) bool) []T {
	firstNoneMatchingIndex := FirstIndexMatching(slice, Negate(predicate))
	if firstNoneMatchingIndex == NotFound {
		return slice
	}
	return slice[firstNoneMatchingIndex:]
}

func DropWhileLast[T any](slice []T, predicate func(T) bool) []T {
	lastNoneMatchingIndex := LastIndexMatching(slice, Negate(predicate))
	if lastNoneMatchingIndex == NotFound {
		return slice
	}
	return slice[0 : lastNoneMatchingIndex+1]
}
