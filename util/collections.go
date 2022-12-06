package util

import (
	"golang.org/x/exp/constraints"
)

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

func Map[T any, R any](slice []T, mapper func(T) R) []R {
	result := make([]R, len(slice))
	for _, element := range slice {
		result = append(result, mapper(element))
	}
	return result
}

func Sum[T constraints.Integer](zeroValue T, slice []T) T {
	result := zeroValue
	for _, element := range slice {
		result += element
	}
	return result
}
