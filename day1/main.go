package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var firstList, secondList []int

	// go through the slice (dynamic array, like list)
	for _, line := range readFile("day1/input") {
		// split by columns (seperated by whitespace)
		var numbers []string = strings.Fields(line)

		// convert the ids to integers
		firstID, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Printf("Could not convert first list item '%s' to integer\n", numbers[0])
		}

		secondID, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Printf("Could not convert second list item '%s' to integer\n", numbers[1])
		}

		firstList = append(firstList, firstID)
		secondList = append(secondList, secondID)
	}

  // part 1
	var differenceSum int = sumDifferenceBetweenSlices(firstList, secondList)
	fmt.Printf("difference between two lists is: %v\n", differenceSum)

  // part 2
	var similarity int = countSimilarity(firstList, secondList)
	fmt.Printf("similarity score between two lists is: %v\n", similarity)
}

func readFile(filepath string) []string {
	inputFile, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("could not read file at '%s'\n", filepath)
		panic("")
	}
	defer inputFile.Close() // remember to close the file afterwards

	// make a scanner that is used to read a file's contents
	scanner := bufio.NewScanner(inputFile)

	// scan the file and split it by each line
	scanner.Split(bufio.ScanLines)

	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text
}

// part 1
func sumDifferenceBetweenSlices(first []int, second []int) int {
	if len(first) != len(second) {
		panic("passed slices are not the same length")
	}

	// sort the lists
	slices.Sort(first)
	slices.Sort(second)

	var sum int = 0
	for index := range first {
		var difference float64 = float64(first[index]) - float64(second[index])
		sum += int(math.Abs(difference))
	}

	return sum
}

// part 2
func countSimilarity(first []int, second []int) int {
  var sum int

  for _, firstID := range first {
    var occurences int
    // get all the occurences of the firstID in the second slice
    for _, secondID := range second {
      if (firstID != secondID) { continue }
      occurences += 1
    }

    sum += occurences * firstID
  }

  return sum
}
