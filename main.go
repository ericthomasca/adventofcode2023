package main

import (
	"fmt"

	"github.com/ericthomasca/adventofcode2023/day01"
	"github.com/ericthomasca/adventofcode2023/day02"
)

func main() {
	fmt.Println("Advent of Code 2023 Solutions")

	day01Part1 := day01.SolvePart1()
	fmt.Println("Day 1, Part 1:", day01Part1)
	
	day01Part2 := day01.SolvePart2()
	fmt.Println("Day 1, Part 2:", day01Part2, "INCORRECT")

	day02Part1 := day02.SolvePart1()
	fmt.Println(day02Part1)

	day02Part2 := day02.SolvePart2()
	fmt.Println(day02Part2)
}