package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	for _, line := range readInputFile() {
		sum += getCalibrationValue(line)
	}
	fmt.Println("Sum:", sum)
}

func getCalibrationValue(text string) int {
	digitNames := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	// ? means unassigned
	leftMost, rightMost := '?', '?'

	assignToCorrectVariable := func(char rune) {
		if leftMost == '?' {
			leftMost, rightMost = char, char
		} else {
			rightMost = char
		}
	}

	for i, char := range text {
		if isDigit(char) {
			assignToCorrectVariable(char)
			continue
		}

		for digitName := range digitNames {
			dnLength := len(digitName)

			// Check if there is enough space for iteration
			if len(text[i:]) < dnLength {
				continue
			}

			if text[i:i+dnLength] == digitName {
				assignToCorrectVariable(digitNames[digitName])
			}
		}
	}

	// sees if the text has no digits at all by checking leftMost
	if leftMost == '?' {
		return 0
	}

	finalString := string(leftMost) + string(rightMost)
	fmt.Println(finalString)

	calibVal, _ := strconv.Atoi(finalString)
	return calibVal
}

func readInputFile() []string {
	contents, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("Error reading the input file")
	}

	lines := strings.Split(string(contents), "\n")

	return lines
}

func isDigit(char rune) bool {
	return '0' <= char && char <= '9'
}
