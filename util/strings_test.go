package util

import (
	"github.com/andrekupka/advent-of-code-2022/assert"
	"testing"
)

func TestIsBlank(t *testing.T) {
	t.Run("should return true for empty string", func(t *testing.T) {
		got := IsBlank("")
		assert.True(t, got)
	})

	t.Run("should return true for string with whitespace", func(t *testing.T) {
		got := IsBlank(" \n\r\t")
		assert.True(t, got)
	})

	t.Run("should return false for string with non-whitespace", func(t *testing.T) {
		got := IsBlank("test")
		assert.False(t, got)
	})
}
