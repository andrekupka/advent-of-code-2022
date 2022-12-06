package util

import (
	"github.com/andrekupka/advent-of-code-2022/test"
	"testing"
)

func TestFirstIndexMatching(t *testing.T) {
	tests := []test.EqualsTest[[]int, int]{
		{
			Name:     "should return -1 if none matches",
			Input:    []int{1, 2, 3},
			Expected: -1,
		},
		{
			Name:     "should return -1 for empty input slice",
			Input:    []int{},
			Expected: -1,
		},
		{
			Name:     "should return index of first matching element",
			Input:    []int{-2, -1, 0, 1},
			Expected: 2,
		},
	}

	isZeroPredicate := func(value int) bool {
		return value == 0
	}

	testFunction := func(input []int) int {
		return FirstIndexMatching(input, isZeroPredicate)
	}

	test.RunEqualsTests(t, tests, testFunction)
}

func TestLastIndexMatching(t *testing.T) {
	tests := []test.EqualsTest[[]int, int]{
		{
			Name:     "should return -1 if none matches",
			Input:    []int{1, 2, 3},
			Expected: -1,
		},
		{
			Name:     "should return -1 for empty input slice",
			Input:    []int{},
			Expected: -1,
		},
		{
			Name:     "should return index of last matching element",
			Input:    []int{-1, 0, 1, 2},
			Expected: 1,
		},
	}

	isZeroPredicate := func(value int) bool {
		return value == 0
	}

	testFunction := func(input []int) int {
		return LastIndexMatching(input, isZeroPredicate)
	}

	test.RunEqualsTests(t, tests, testFunction)
}

func TestDropWhile(t *testing.T) {
	tests := []test.EqualsTest[[]int, []int]{
		{
			Name:     "should keep empty slice",
			Input:    []int{},
			Expected: []int{},
		},
		{
			Name:     "should not drop if condition does not match",
			Input:    []int{0, 1, 2, 3},
			Expected: []int{0, 1, 2, 3},
		},
		{
			Name:     "should drop from start until element does not match predicate",
			Input:    []int{-2, -1, 0, 1},
			Expected: []int{0, 1},
		},
	}

	isNegativePredicate := func(value int) bool {
		return value < 0
	}

	testFunction := func(input []int) []int {
		return DropWhile(input, isNegativePredicate)
	}

	test.RunEqualsTests(t, tests, testFunction)
}

func TestDropWhileLast(t *testing.T) {
	tests := []test.EqualsTest[[]int, []int]{
		{
			Name:     "should keep empty slice",
			Input:    []int{},
			Expected: []int{},
		},
		{
			Name:     "should not drop if condition does not match",
			Input:    []int{0, 1, 2, 3},
			Expected: []int{0, 1, 2, 3},
		},
		{
			Name:     "should drop from end until element does not match predicate",
			Input:    []int{2, 1, 0, -1, -2},
			Expected: []int{2, 1, 0},
		},
	}

	isNegativePredicate := func(value int) bool {
		return value < 0
	}

	testFunction := func(input []int) []int {
		return DropWhileLast(input, isNegativePredicate)
	}

	test.RunEqualsTests(t, tests, testFunction)
}
