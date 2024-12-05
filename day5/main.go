package main

import (
	"fmt"
	"sakuexe/adventofcode2024/utils"
	"slices"
	"strings"
)

func splitOrdersFromUpdates(input []string) (map[int][]int, [][]int) {
  orders := map[int][]int{}
  updates := [][]int{}

  isOrder := true
  for _, line := range input {
    // look for the seperator (double new line without the newlines)
    if line == "" {
      isOrder = false
      continue
    }

    // while the sperator is not found, parse as orders
    if isOrder {
      values := strings.Split(line, "|")
      firstValue, secondValue := utils.ParseInt(values[0]), utils.ParseInt(values[1])
      orders[firstValue] = append(orders[firstValue], secondValue)
      continue
    }

    // parse the rest as updates
    values := strings.Split(line, ",")
    updatePages := []int{}
    for _, value := range values {
      updatePages = append(updatePages, utils.ParseInt(value))
    }

    updates = append(updates, updatePages)
  }

  return orders, updates
}

func main() {
  inputValue := utils.ReadFile("day5/input")

  orders, updates := splitOrdersFromUpdates(inputValue)

  result := getValidUpdates(orders, updates)
  fmt.Println(result)
}

func getValidUpdates(orders map[int][]int, updates [][]int) int {
  score := 0

  for _, update := range updates {
    seenValues := []int{}
    isInvalid := false

    for _, page := range update {
      seenValues = append(seenValues, page)

      // update does not have requirements
      if orders[page] == nil {
        continue
      }

      // fmt.Println("checking for:", update, "dependencies")
      // fmt.Println("values seen are:", seenValues)

      for _, lastValue := range orders[page] {
        if slices.Contains(seenValues, lastValue) {
          // fmt.Println(lastValue, "[INVALID] was seen before")
          isInvalid = true;
          break
        }
      }
    }

    if isInvalid {
      continue
    }
    score += update[len(update) / 2]
  }

  return score
}
