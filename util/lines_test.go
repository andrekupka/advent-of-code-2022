package util

import (
	"github.com/andrekupka/advent-of-code-2022/test"
	"testing"
)

func TestLinesFromString(t *testing.T) {
	tests := []test.EqualsTest[string, *Lines]{
		{
			Name:     "should return single empty line for empty input",
			Input:    "",
			Expected: &Lines{lines: []string{""}},
		},
		{
			Name:     "should return all lines",
			Input:    "first\nsecond\nlast",
			Expected: &Lines{lines: []string{"first", "second", "last"}},
		},
		{
			Name:     "should keep blank lines",
			Input:    "  \nfirst\n \nsecond\n",
			Expected: &Lines{lines: []string{"  ", "first", " ", "second", ""}},
		},
	}

	test.RunEqualsTests(t, tests, LinesFromString)
}

func TestLines_TrimBlankLines(t *testing.T) {
	tests := []test.EqualsTest[*Lines, *Lines]{
		{
			Name:     "should keep lines as is if none are blank",
			Input:    &Lines{lines: []string{"first", "second"}},
			Expected: &Lines{lines: []string{"first", "second"}},
		},
		{
			Name:     "should trim blank lines from both sides",
			Input:    &Lines{lines: []string{" ", "first", "second", "", "\n"}},
			Expected: &Lines{lines: []string{"first", "second"}},
		},
	}

	test.RunEqualsTests(t, tests, func(input *Lines) *Lines {
		return input.TrimBlankLines()
	})
}
