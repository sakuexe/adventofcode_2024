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

func convertReport(strSlice []string) []int {
	var convertedNumbers []int

	for _, number := range strSlice {
		convertedNumber, err := strconv.Atoi(number)
		if err != nil {
			panic("string '" + number + "' could not be converted to int")
		}
		convertedNumbers = append(convertedNumbers, convertedNumber)
	}

	return convertedNumbers
}

func main() {
	var inputText []string = readFile("day2/input")
	var reports [][]int
  var safeReports int

  // get all the reports to int slices
	for _, line := range inputText {
		var rawNumbers []string = strings.Fields(line)
		reports = append(reports, convertReport(rawNumbers))
	}

  // validate the reports (part 1)
  // make a copy of the reports, to pass to the second part
  // because golang passes slices as references by default
  for _, report := range slices.Clone(reports) {
    if (isReportSafe(report)) {
      safeReports += 1
    }
  }
  fmt.Printf("Safe reports: %v\n", safeReports)

  safeReports = 0
  // include damper of 1 (part 2)
  for _, report := range slices.Clone(reports) {
    if (checkReportWithDampening(report)) {
      safeReports += 1
    }
  }
  fmt.Printf("Safe reports (dampened): %v\n", safeReports)
}

func isReportSafe(report []int) bool {
  var previousLevel int
  var isIncreasing bool

  for index, level := range report {
    // initialize
    if (index == 0) {
      previousLevel = level
      continue
    }
    // check if the values are increasing or decreasing
    if (index == 1) {
      isIncreasing = level > previousLevel
    }

    // is the increase/decrease continuing
    if (isIncreasing && previousLevel > level) {
      return false
    }
    if (!isIncreasing && previousLevel < level) {
      return false
    }

    var difference int = int(math.Abs(float64(level - previousLevel)))

    // check for valid change range
    if (difference < 1 || difference > 3) {
      return false
    }

    // remember to update the previous level (I forgor and was confused)
    previousLevel = level
  }
	return true
}

func checkReportWithDampening(report []int) bool {
  if (isReportSafe(report)) {
    return true
  }

  // try removing one level and checking if the report becomes safe
  for index, _ := range report {
    var reportWithoutCurrentValue []int = removeIndexFromSlice(report, index)
    if (isReportSafe(reportWithoutCurrentValue)) {
      return true
    }
  }
  return false
}

func removeIndexFromSlice(slice []int, index int) []int {
    reportWithoutCurrentValue := make([]int, 0)
    reportWithoutCurrentValue = append(reportWithoutCurrentValue, slice[:index]...)
    return append(reportWithoutCurrentValue, slice[index+1:]...)
}
