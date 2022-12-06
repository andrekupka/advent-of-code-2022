package test

import (
	"github.com/andrekupka/advent-of-code-2022/assert"
	"testing"
)

type EqualsTest[I any, E any] struct {
	Name     string
	Input    I
	Expected E
}

func NewEqualsTest[I any, E any](name string, input I, expected E) *EqualsTest[I, E] {
	return &EqualsTest[I, E]{
		Name:     name,
		Input:    input,
		Expected: expected,
	}
}

func RunEqualsTests[I any, E any](t *testing.T, testData []EqualsTest[I, E], testFunction func(I) E) {
	for _, test := range testData {
		t.Run(test.Name, func(t *testing.T) {
			actual := testFunction(test.Input)
			assert.Equals(t, test.Expected, actual)
		})
	}
}
