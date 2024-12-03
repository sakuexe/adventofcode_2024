package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func main() {
	textInput := readFile("day3/input")
	instructions := getValidInstructions(textInput)
  // part 1
	result := applyInstructions(instructions)
	fmt.Println(result)
}

func getValidInstructions(memoryValues []string) [][]string {
	validInstructions := [][]string{}
	for _, memory := range memoryValues {
		filterExpression, err := regexp.Compile(`mul\(\d+\,\d+\)`)
		if err != nil {
			panic("error trying to compile regular expression")
		}

		matches := filterExpression.FindAllStringSubmatch(memory, -1)
		if matches == nil {
			panic("did not find any matches for regex in input")
		}
    validInstructions = append(validInstructions, matches...)
	}
	return validInstructions
}

func applyInstructions(instructions [][]string) int {
	var sum int

	for _, match := range instructions {
		// prepare the expression
		numberExpression, err := regexp.Compile(`(?P<FirstValue>\d+)\,(?P<SecondValue>\d+)`)
		if err != nil {
			panic("error trying to compile regular expression")
		}

		// get the numbers being applied
		parameters := numberExpression.FindStringSubmatch(match[0])
		firstGroup, secondGroup := numberExpression.SubexpIndex("FirstValue"), numberExpression.SubexpIndex("SecondValue")

		// convert string values to int
		firstValue, err := strconv.Atoi(parameters[firstGroup])
		if err != nil {
			panic("could not convert '" + parameters[firstGroup] + "' to int")
		}
		secondValue, err := strconv.Atoi(parameters[secondGroup])
		if err != nil {
			panic("could not convert '" + parameters[secondGroup] + "' to int")
		}

		// apply the multiplication
		// fmt.Printf("multiplying %d * %d\n", firstValue, secondValue)
		sum += firstValue * secondValue
	}
	return sum
}
