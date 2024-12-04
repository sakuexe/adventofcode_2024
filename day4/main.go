package main

import (
	"fmt"
	"sakuexe/adventofcode2024/utils"
	"strconv"
)

var solutionMap = make(map[int]string)

func main() {
  solutionMap[0] = "X"
  solutionMap[1] = "M"
  solutionMap[2] = "A"
  solutionMap[3] = "S"

  inputText := utils.ReadFile("day4/input.test")

  fmt.Println("XMAS")
  fmt.Printf("found %d hits in row\n", checkRow(inputText[0]))
}

func solveSearch(wordSearch []string) {
}

func checkRow(row string) int {
  foundInRow := 0

  forward := ""
  for _, character := range row {
    fmt.Printf("current character is: %c\n", character)
    fmt.Printf("wanted character: %s\n", solutionMap[len(forward)])
    fmt.Println("forward value: ", forward)
    currentChar := strconv.QuoteRune(character)
    if solutionMap[len(forward)] != currentChar {
      fmt.Println("[chars do not match]")
      forward = ""
      continue
    }
    forward += currentChar
    if forward == "XMAS" {
      foundInRow += 1
    }
  }

  return foundInRow
}
