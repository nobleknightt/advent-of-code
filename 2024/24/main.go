package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hasValue(wire string, wireValues map[string]int) bool {
	var _, ok = wireValues[wire]
	return ok
}

func calculateOutput(operation string, valueOne, valueTwo int) int {
	if operation == "AND" {
		return valueOne & valueTwo
	} else if operation == "OR" {
		return valueOne | valueTwo
	} else {
		return valueOne ^ valueTwo
	}
}

func solvePartOne(input []string) int {
	var wireValues = make(map[string]int)
	var gates [][]string
	for _, line := range input {
		var splits = strings.Fields(line)
		if len(splits) == 2 {
			var key = strings.TrimRightFunc(splits[0], func(r rune) bool { return r == ':' })
			var value, _ = strconv.Atoi(splits[1])
			wireValues[key] = value
		} else {
			gates = append(gates, []string{splits[1], splits[0], splits[2], splits[4]})
		}
	}

	for len(gates) > 0 {
		var gate = gates[0]
		gates = gates[1:]
		if hasValue(gate[1], wireValues) && hasValue(gate[2], wireValues) {
			wireValues[gate[3]] = calculateOutput(gate[0], wireValues[gate[1]], wireValues[gate[2]])
		} else {
			gates = append(gates, gate)
		}
	}

	var output = 0
	for key, value := range wireValues {
		if key[0] == 'z' {
			var shift, _ = strconv.Atoi(key[1:])
			output = output | (value << shift) 
		}
	}

	return output
}

func solvePartTwo(input []string) int {
	return len(input)
}

func main() {
	var input []string

	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line = scanner.Text()
		if len(line) > 0 {
			input = append(input, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return
	}

	var partOneAnswer = solvePartOne(input)
	fmt.Println(partOneAnswer)

	var partTwoAnswer = solvePartTwo(input)
	fmt.Println(partTwoAnswer)
}
