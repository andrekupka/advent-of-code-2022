package day01

import (
	"fmt"
	"github.com/andrekupka/advent-of-code-2022/util"
	"strconv"
)

func Solve(input string) {
	linesGroups := util.LinesFromString(input).GroupByConsecutiveNonBlankLines()

	elfIndex := 0
	maximumCalories := 0

	for index, linesGroup := range linesGroups {
		numbers := util.Map(linesGroup.Value(), func(line string) int {
			result, err := strconv.Atoi(line)
			if err != nil {
				panic("Expected number as input")
			}
			return result
		})
		calories := util.Sum(0, numbers)

		if calories > maximumCalories {
			maximumCalories = calories
			elfIndex = index
		}
	}

	fmt.Printf("Elf number %d carries the most (%d calories)!", (elfIndex + 1), maximumCalories)
}
