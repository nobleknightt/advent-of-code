package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func solvePartOne(input []string) int {
	var reInstruction = regexp.MustCompile(`mul\([0-9]+\,[0-9]+\)`)
	var reNumber = regexp.MustCompile(`[0-9]+`)

	var result = 0
	for _, line := range input {
		var instructionMatches = reInstruction.FindAllString(line, -1)

		for _, value := range instructionMatches {
			var numberMatches = reNumber.FindAllString(value, -1)
			var a, _ = strconv.Atoi(numberMatches[0])
			var b, _ = strconv.Atoi(numberMatches[1])
			result += a * b
		}
	}

	return result
}

func solvePartTwo(input []string) int {
	var reInstruction = regexp.MustCompile(`mul\([0-9]+\,[0-9]+\)|do\(\)|don't\(\)`)
	var reNumber = regexp.MustCompile(`[0-9]+`)

	var result = 0

	var enabled = true
	for _, line := range input {
		var instructionMatches = reInstruction.FindAllString(line, -1)
		for _, value := range instructionMatches {
			if value == "do()" {
				enabled = true
			} else if value == "don't()" {
				enabled = false
			} else if enabled {
				var numberMatches = reNumber.FindAllString(value, -1)
				var a, _ = strconv.Atoi(numberMatches[0])
				var b, _ = strconv.Atoi(numberMatches[1])
				result += a * b
			}
		}
	}

	return result
}

func main() {
	var input []string

	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line = scanner.Text()
		input = append(input, line)
	}

	var partOneAnswer = solvePartOne(input)
	fmt.Println(partOneAnswer)

	var partTwoAnswer = solvePartTwo(input)
	fmt.Println(partTwoAnswer)
}
