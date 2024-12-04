package main

import (
	"fmt"
	"sakuexe/adventofcode2024/utils"
	"slices"
)

var solutionMap = make(map[int]rune)

func stringToMatrix(stringSlice []string) [][]rune {
  matrix := [][]rune{}

  for _, row := range stringSlice {
    matrix = append(matrix, []rune(row))
  }

  return matrix
}

func main() {
	solutionMap[0] = 'X'
	solutionMap[1] = 'M'
	solutionMap[2] = 'A'
	solutionMap[3] = 'S'

	inputText := utils.ReadFile("day4/input.test")
  wordSearchMatrix := stringToMatrix(inputText)

	result := solveSearch(wordSearchMatrix)

	fmt.Printf("found %d hits in row\n", result)
}

func solveSearch(wordSearch [][]rune) int {
	totalScore := 0

  // search for rows
	for index, row := range wordSearch {

    // row
		totalScore += checkRow(row)
		if checkRow(row) != 0 {
			fmt.Printf("line: #%d has a hit\n", index)
		}

    // column
    totalScore += checkColumn(index, wordSearch, false)
		if checkColumn(index, wordSearch, false) != 0 {
			fmt.Printf("column: #%d has a hit\n", index)
		}

    reversedRow := slices.Clone(row)
    slices.Reverse(reversedRow)

    // reversed row
		totalScore += checkRow(reversedRow)
		if checkRow(reversedRow) != 0 {
			fmt.Printf("line: #%d has a reverse hit\n", index)
		}

    // reversed column
		totalScore += checkColumn(index, wordSearch, true)
		if checkColumn(index, wordSearch, true) != 0 {
			fmt.Printf("column: #%d has a reverse hit\n", index)
		}
	}

  totalScore += checkDiagonal(wordSearch, false)
  totalScore += checkDiagonal(wordSearch, true)
  slices.Reverse(wordSearch)
  for _, row := range wordSearch {
    slices.Reverse(row)
  }
  totalScore += checkDiagonal(wordSearch, false)
  totalScore += checkDiagonal(wordSearch, true)
  
  // search for columns
	return totalScore
}

func checkRow(row []rune) int {
	foundInRow := 0
	forward := []rune{}

	for _, character := range row {

    // if the characters do not follow the wanted solution
		if solutionMap[len(forward)] != character {
			// reset the count, unless its an x
			if character == solutionMap[0] {
				forward = []rune{character}
			} else {
				forward = []rune{}
			}
			continue
		}

    // the next character is valid
		forward = append(forward, character)
		if len(forward) == 4 {
			foundInRow += 1
			forward = []rune{}
		}

	}

	return foundInRow
}

func checkColumn(columnIndex int, matrix [][]rune, reverse bool) int {
  foundInColumn := 0
	forward := []rune{}

  for index, _ := range matrix {

    currentRow := matrix[index]
    // check in reverse if reverse is passed as true
    if reverse {
      currentRow = matrix[len(matrix) - 1 - index]
    }

    character := currentRow[columnIndex]

    // if the characters do not follow the wanted solution
		if solutionMap[len(forward)] != character {
			// reset the count, unless its an x
			if character == solutionMap[0] {
				forward = []rune{character}
			} else {
				forward = []rune{}
			}
			continue
		}

    // the next character is valid
		forward = append(forward, character)
		if len(forward) == 4 {
			foundInColumn += 1
			forward = []rune{}
		}
  }

  return foundInColumn
}

func checkDiagonal(matrix [][]rune, reversed bool) int {
  rows := len(matrix)
  cols := len(matrix[0])
  forward := []rune{}
  foundDiagonally := 0

  // walk through the matrix diagonally
  // https://www.geeksforgeeks.org/zigzag-or-diagonal-traversal-of-matrix/
  for line := range rows + cols {

    start_col := max(0, line - rows)
    elements := min(line, (cols - start_col), rows)

    for j := range elements {
      character := matrix[min(rows, line) - j - 1][start_col + j]
      if reversed {
        character = matrix[start_col + j][min(rows, line) - j - 1]
        fmt.Println(string(character))
      }

      // if the characters do not follow the wanted solution
      if solutionMap[len(forward)] != character {
        // reset the count, unless its an x
        if character == solutionMap[0] {
          forward = []rune{character}
        } else {
          forward = []rune{}
        }
        continue
      }

      // the next character is valid
      forward = append(forward, character)
      if len(forward) == 4 {
        foundDiagonally += 1
        forward = []rune{}
      }
    }
  }

  return foundDiagonally
}
