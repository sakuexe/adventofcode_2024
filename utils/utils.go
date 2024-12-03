package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ParseInt(str string) int {
  convertedString, err := strconv.Atoi(str)
  if err != nil {
    panic("could not convert '" + str + "' to int")
  }
  return convertedString
}

func ReadFile(filepath string) []string {
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
