package main

import (
	"fmt"
	"github.com/andrekupka/advent-of-code-2022/day01"
	"github.com/andrekupka/advent-of-code-2022/util"
	"os"
)

func main() {
	println("Hello Advent of Code 2022!")

	path := "inputs/day01_1.txt"

	input, err := util.ReadFileAsString(path)
	if err != nil {
		fmt.Println("Failed to load input from", path)
		os.Exit(1)
	}

	day01.Solve(input)
}
