package day01

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	// "strings"
)

func SolvePart1() int {
	file, err := os.Open("day01/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var calibrationValueTotal = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		nonDigitRegexp := regexp.MustCompile(`\D`)
		digitsOnly := nonDigitRegexp.ReplaceAllString(text, "")

		digits := []rune(digitsOnly)
		
		firstDigit := int(digits[0] - '0')
		lastDigit := int(digits[len(digits) - 1] - '0')

		calibrationValue := (firstDigit * 10) + lastDigit

		calibrationValueTotal += calibrationValue
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return calibrationValueTotal
}

func SolvePart2() int {
	file, err := os.Open("day01/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var calibrationValueTotal = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		// convert numbers as words to numbers as digits
		
		nonDigitRegexp := regexp.MustCompile(`\D`)
		digitsOnly := nonDigitRegexp.ReplaceAllString(text, "")

		digits := []rune(digitsOnly)
		
		firstDigit := int(digits[0] - '0')
		lastDigit := int(digits[len(digits) - 1] - '0')

		calibrationValue := (firstDigit * 10) + lastDigit

		calibrationValueTotal += calibrationValue
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return calibrationValueTotal
}
