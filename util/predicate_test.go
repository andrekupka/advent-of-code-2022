package util

import (
	"github.com/andrekupka/advent-of-code-2022/assert"
	"testing"
)

func TestNegate(t *testing.T) {

	t.Run("should create predicate that returns negated result", func(t *testing.T) {
		isZeroPredicate := func(value int) bool {
			return value == 0
		}

		negatedPredicate := Negate(isZeroPredicate)

		assert.Equals(t, !isZeroPredicate(0), negatedPredicate(0))
	})
}
