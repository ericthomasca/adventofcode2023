package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SolvePart1() int {
	file, err := os.Open("day02/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	idSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		fullGameId := strings.Split(text, ":")[0]
		gameIdString := strings.Split(fullGameId, " ")[1]
		gameId, err := strconv.Atoi(gameIdString)
		if err != nil {
			fmt.Println(err)
		}

		handsString := strings.Split(text, ":")[1]
		handsString = strings.TrimSpace(handsString)
		hands := strings.Split(handsString, ";")
		
		handsPossible := []bool{}
		
		for _, hand := range hands {
			hand = strings.TrimSpace(hand)
			
			colorGroups := strings.Split(hand, ",")
			
			for _, colorGroup := range colorGroups {
				colorGroup = strings.TrimSpace(colorGroup)
				cubeColor := strings.Split(colorGroup, " ")[1]
				
				cubeAmountString := (strings.Split(colorGroup, " ")[0])
				cubeAmount, err := strconv.Atoi(cubeAmountString)
				if err != nil {
					fmt.Println(err)
				}
				
				isPossibleHand := isPossibleHand(cubeAmount, cubeColor)
				handsPossible = append(handsPossible, isPossibleHand)
			}
		}
		
		fmt.Println(gameId)
		
		gamePossible := true
		for _, hand := range handsPossible {
			if !hand {
				gamePossible = false
				break
			}
		}

		fmt.Println(gamePossible)

		if gamePossible {
			idSum += gameId
		}
	}

	return idSum
}

func isPossibleHand(cubeAmount int, cubeColor string) bool {
	switch {
	case cubeAmount <= 12 && cubeColor == "red":
		return true
	case cubeAmount <= 13 && cubeColor == "green":
		return true
	case cubeAmount <= 14 && cubeColor == "blue":
		return true
	default:
		return false 
	}
}
