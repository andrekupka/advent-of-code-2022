package util

import "strings"

type Lines struct {
	lines []string
}

func NewLines(lines []string) *Lines {
	return &Lines{lines: lines}
}

func LinesFromString(input string) *Lines {
	lines := strings.Split(input, "\n")
	return NewLines(lines)
}

func (l *Lines) TrimBlankLines() *Lines {
	intermediateLines := DropWhile(l.lines, IsBlank)
	trimmedLines := DropWhileLast(intermediateLines, IsBlank)
	return &Lines{lines: trimmedLines}
}
