package main

import (
	"fmt"
	"reflect"
	"sakuexe/adventofcode2024/utils"
	"slices"
)

var (
	solutionSlice         []rune
	reversedSolutionSlice []rune
)

func stringToMatrix(stringSlice []string) [][]rune {
	matrix := [][]rune{}

	for _, row := range stringSlice {
		matrix = append(matrix, []rune(row))
	}

	return matrix
}

func main() {
  // setup the validation variables
	solutionSlice = []rune("XMAS")
	reversedSolutionSlice = slices.Clone(solutionSlice)
	slices.Reverse(reversedSolutionSlice)

  // get a matrix of runes (character codes)
	inputText := utils.ReadFile("day4/input")
	wordSearchMatrix := stringToMatrix(inputText)

	// part 1
	result := solveSearch(wordSearchMatrix)
	fmt.Printf("found %d hits in row\n", result)

  // part 2 (maybe one day)
}

func printMatrix(matrix [][]rune) {
	for _, row := range matrix {
		column := []string{}
		for _, col := range row {
			column = append(column, string(col))
		}
		fmt.Println(column)
	}
}

func solveSearch(wordSearch [][]rune) int {
	totalScore := 0
	rotations := [][][]rune{
		wordSearch,
		utils.RotateMatrix45Deg(wordSearch),
		utils.RotateMatrix90Deg(wordSearch),
		utils.RotateMatrix45Deg(utils.RotateMatrix90Deg(wordSearch)),
	}

	for index, rotation := range rotations {
		for row := range rotation {
			for col := range rotation[row] {
				totalScore += checkHorizontal(row, col, rotation)
			}
		}
		fmt.Println("===ROTATION:", index, "===")
		printMatrix(rotation)
	}

	return totalScore
}

func checkHorizontal(y int, x int, matrix [][]rune) int {
	if len(matrix[y][x:]) < 4 {
		return 0
	}

	// forward (XMAS)
	if reflect.DeepEqual(solutionSlice, matrix[y][x:x+4]) {
		return 1
	}
	// reverse (SAMX)
	if reflect.DeepEqual(reversedSolutionSlice, matrix[y][x:x+4]) {
		return 1
	}

	return 0
}
