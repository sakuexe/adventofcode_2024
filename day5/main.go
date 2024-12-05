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
  inputValue := utils.ReadFile("day5/input.test")

  orders, updates := splitOrdersFromUpdates(inputValue)

  // part 1
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
      fmt.Println("old order:", update)
      fmt.Println("new order:", reorderUpdate(orders, update))
      continue
    }
    score += update[len(update) / 2]
  }

  return score
}

func reorderUpdate(orders map[int][]int, update []int) []int {
  // keeps track of the value and it's index inside the update
  seenValues := map[int]int{}

  for index, page := range update {
    seenValues[page] = index
    // update does not have requirements
    if orders[page] == nil {
      continue
    }

    for _, lastValue := range orders[page] {
      // if the requirement is not seen
      if _, ok := seenValues[lastValue]; !ok {
        continue
      }
      requirementIndex := seenValues[lastValue]
      beforeValue := update[requirementIndex]

      // get the values before the value that we want to get in front of
      newUpdateOrder := append(update[:requirementIndex], page)
      // place the old value after the current one
      newUpdateOrder = append(newUpdateOrder, beforeValue)
      // fill in the rest
      newUpdateOrder = append(newUpdateOrder, update[index+1:]...)
      // re-validate the order
      reorderUpdate(orders, newUpdateOrder)
    }
  }
  return update
}
