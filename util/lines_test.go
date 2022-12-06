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

func TestLines_Append(t *testing.T) {
	line := "appendedLine"

	tests := []test.EqualsTest[*Lines, *Lines]{
		{
			Name:     "should append to empty lines",
			Input:    &Lines{lines: []string{}},
			Expected: &Lines{lines: []string{line}},
		},
		{
			Name:     "should append after last line",
			Input:    &Lines{lines: []string{"first", "last"}},
			Expected: &Lines{lines: []string{"first", "last", line}},
		},
	}

	testFunction := func(input *Lines) *Lines {
		return input.Append(line)
	}

	test.RunEqualsTests(t, tests, testFunction)
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

func TestLines_GroupByConsecutiveNonBlankLines(t *testing.T) {
	tests := []test.EqualsTest[*Lines, []Lines]{
		{
			Name:     "should return no groups from no lines",
			Input:    &Lines{lines: []string{}},
			Expected: []Lines{},
		},
		{
			Name:     "should return no groups from only blank lines",
			Input:    &Lines{lines: []string{"", "  "}},
			Expected: []Lines{},
		},
		{
			Name:  "should return a single group from consecutive lines",
			Input: &Lines{lines: []string{"a", "b"}},
			Expected: []Lines{
				{lines: []string{"a", "b"}},
			},
		},
		{
			Name:  "should return multiple groups if separated by a blank line",
			Input: &Lines{lines: []string{"a", " ", "b", "c"}},
			Expected: []Lines{
				{lines: []string{"a"}},
				{lines: []string{"b", "c"}},
			},
		},
		{
			Name:  "should skip multiple blank lines without creating new groups",
			Input: &Lines{lines: []string{"a", " ", " ", "b"}},
			Expected: []Lines{
				{lines: []string{"a"}},
				{lines: []string{"b"}},
			},
		},
		{
			Name:  "should not create empty groups for surrounding blank lines",
			Input: &Lines{lines: []string{" ", "a", " "}},
			Expected: []Lines{
				{lines: []string{"a"}},
			},
		},
	}

	testFunction := func(l *Lines) []Lines {
		return l.GroupByConsecutiveNonBlankLines()
	}

	test.RunEqualsTests(t, tests, testFunction)
}
