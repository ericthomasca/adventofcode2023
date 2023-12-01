package day01

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func SolveDay1() {
	file, err := os.Open("day01/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		nonDigitRegexp := regexp.MustCompile(`\D`)
		digitsOnly := nonDigitRegexp.ReplaceAllString(text, "")

		fmt.Println(digitsOnly)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}