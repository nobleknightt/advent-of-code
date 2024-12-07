package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func operatorAdd(a, b int) int {
	return a + b
}

func operatorMultiply(a, b int) int {
	return a * b
}

func operatorConcatenate(a, b int) int {
	var c, _ = strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return c
}

func checkIfEquationTrue(i, result, requiredResult int, numbers []int, operators []func(int, int) int) bool {
	if i == len(numbers) {
		return result == requiredResult
	}

	for _, operator := range operators {
		var isEquationTrue = checkIfEquationTrue(i+1, operator(result, numbers[i]), requiredResult, numbers, operators)
		if isEquationTrue {
			return true
		}
	}

	return false
}

func solvePartOne(input [][]int) int {
	var total = 0

	var operators = []func(int, int) int{operatorAdd, operatorMultiply}

	for i := 0; i < len(input); i++ {
		var isEquationTrue = checkIfEquationTrue(1, input[i][1], input[i][0], input[i][1:], operators)
		if isEquationTrue {
			total += input[i][0]
		}
	}

	return total
}

func solvePartTwo(input [][]int) int {
	var total = 0

	var operators = []func(int, int) int{operatorAdd, operatorMultiply, operatorConcatenate}

	for i := 0; i < len(input); i++ {
		var isEquationTrue = checkIfEquationTrue(1, input[i][1], input[i][0], input[i][1:], operators)
		if isEquationTrue {
			total += input[i][0]
		}
	}

	return total
}

func main() {
	var input [][]int

	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line = scanner.Text()

		var values []int
		for i, value := range strings.Fields(line) {
			if i == 0 {
				value = strings.TrimSuffix(value, ":")
			}
			var valueAsInt, _ = strconv.Atoi(value)
			values = append(values, valueAsInt)
		}
		input = append(input, values)
	}

	var partOneAnswer = solvePartOne(input)
	fmt.Println(partOneAnswer)

	var partTwoAnswer = solvePartTwo(input)
	fmt.Println(partTwoAnswer)
}
