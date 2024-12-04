package main

import (
	"fmt"
	"sakuexe/adventofcode2024/utils"
)

var solutionMap = make(map[int]rune)

func main() {
	solutionMap[0] = 'X'
	solutionMap[1] = 'M'
	solutionMap[2] = 'A'
	solutionMap[3] = 'S'

	inputText := utils.ReadFile("day4/input.test")

	result := solveSearch(inputText)
	fmt.Printf("found %d hits in row\n", result)
}

func solveSearch(wordSearch []string) int {
	totalScore := 0
	for index, row := range wordSearch {
		totalScore += checkRow(row)
		if checkRow(row) != 0 {
			fmt.Printf("line: #%d has a hit\n", index)
		}
	}
	return totalScore
}

func checkRow(row string) int {
	foundInRow := 0
	forward := []rune{}
	backward := []rune{}

	for _, character := range row {

		// handle the forward search
		if solutionMap[len(forward)] != character {
			// reset the count, unless its an x
			if character == 'X' {
				forward = []rune{'X'}
			} else {
				forward = []rune{}
			}
		} else {
			forward = append(forward, character)
			if len(forward) == 4 {
				foundInRow += 1
				forward = []rune{}
			}
		}

		// handle the reverse search
		if solutionMap[3 - len(backward)] != character {
			// reset the count, unless its an x
			if character == 'S' {
				backward = []rune{'S'}
			} else {
				backward = []rune{}
			}
		} else {
			backward = append(backward, character)
			if len(backward) == 4 {
				foundInRow += 1
				backward = []rune{}
			}
		}
	}

	return foundInRow
}
