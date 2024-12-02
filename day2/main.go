package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func test(slice []int) {
  slice[0] = 200
}

func main() {
	var inputText []string = readFile("day2/input.test")
	var reports [][]int
  var safeReports int

  // get all the reports to int slices
	for _, line := range inputText {
		var rawNumbers []string = strings.Fields(line)
		reports = append(reports, convertReport(rawNumbers))
	}

  // make a copy of the reports, to pass to the second part
  // because golang passes slices as references by default
  reportsCopy := make([][]int, len(reports))
  for i := range reports {
      reportsCopy[i] = make([]int, len(reports[i]))
      copy(reportsCopy[i], reports[i])
  }
  fmt.Println(reportsCopy)

  // validate the reports (part 1)
  for _, report := range reports {
    if (!isReportSafe(report, 0)) {
      continue
    }
    safeReports += 1
  }
  fmt.Printf("Safe reports: %v\n", safeReports)

  safeReports = 0
  // include damper of 1 (part 2)
  for index, report := range reportsCopy {
    fmt.Println("=== Report #", index, "===")
    if (!isReportSafe(report, 1)) {
      continue
    }
    safeReports += 1
  }
  fmt.Printf("Safe reports (dampened): %v\n", safeReports)
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

func isReportSafe(report []int, damperValue int) bool {
  fmt.Println("report:", report)
  var previousLevel int
  var isIncreasing bool
  if (damperValue < 0) {
    fmt.Println("damper limit reached")
    return false
  }

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
      // fmt.Println("increase did not continue")
      // fmt.Println("level:", level, "previous:", previousLevel)
      var reportWithoutCurrentValue []int = append(report[:index - 1], report[index:]...)
      return isReportSafe(reportWithoutCurrentValue, damperValue - 1)
    }
    if (!isIncreasing && previousLevel < level) {
      // fmt.Println("decrease did not continue")
      // fmt.Println("level:", level, "previous:", previousLevel)
      var reportWithoutCurrentValue []int = append(report[:index - 1], report[index:]...)
      return isReportSafe(reportWithoutCurrentValue, damperValue - 1)
    }

    var difference int = int(math.Abs(float64(level - previousLevel)))

    // check for valid change range
    if (difference < 1 || difference > 3) {
      // fmt.Printf("change rate is not valid: abs(%v - %v) = %v\n", level, previousLevel, difference)
      // fmt.Println("level:", level, "previous:", previousLevel)
      var reportWithoutCurrentValue []int = append(report[:index - 1], report[index:]...)
      return isReportSafe(reportWithoutCurrentValue, damperValue - 1)
    }

    // remember to update the previous level (I forgor and was confused)
    previousLevel = level
  }
	return true
}
