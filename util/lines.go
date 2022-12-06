package util

import "strings"

type Lines struct {
	lines []string
}

func NewLines(lines []string) *Lines {
	return &Lines{lines: lines}
}

func EmptyLines() *Lines {
	return &Lines{[]string{}}
}

func (l *Lines) Value() []string {
	return l.lines
}

func LinesFromString(input string) *Lines {
	lines := strings.Split(input, "\n")
	return NewLines(lines)
}

func (l *Lines) Append(line string) *Lines {
	return &Lines{lines: append(l.lines, line)}
}

func (l *Lines) TrimBlankLines() *Lines {
	intermediateLines := DropWhile(l.lines, IsBlank)
	trimmedLines := DropWhileLast(intermediateLines, IsBlank)
	return &Lines{lines: trimmedLines}
}

func (l *Lines) GroupByConsecutiveNonBlankLines() []Lines {
	startNewGroup := true

	linesGroups := make([]Lines, 0)

	var currentLineGroup *Lines

	for _, line := range l.lines {
		if IsBlank(line) {
			if currentLineGroup != nil {
				linesGroups = append(linesGroups, *currentLineGroup)
				currentLineGroup = nil
				startNewGroup = true
			}
		} else {
			if startNewGroup {
				startNewGroup = false
				currentLineGroup = &Lines{lines: []string{line}}
			} else {
				currentLineGroup = currentLineGroup.Append(line)
			}
		}
	}

	if currentLineGroup != nil {
		linesGroups = append(linesGroups, *currentLineGroup)
	}

	return linesGroups
}
