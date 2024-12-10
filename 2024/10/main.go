package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findReachablePositions(i, j int, mapOfArea [][]int, shouldFindUnique bool) int {

	var reachablePositions [][]int
	var stack = [][]int{{i, j}}
	for len(stack) > 0 {
		var position = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		var i, j = position[0], position[1]
		if mapOfArea[i][j] == 9 {
			reachablePositions = append(reachablePositions, []int{i, j})
		}
		if i + 1 < len(mapOfArea) && mapOfArea[i + 1][j] == mapOfArea[i][j] + 1 {
			stack = append(stack, []int{i + 1, j})
		}
		if i - 1 >= 0 && mapOfArea[i - 1][j] == mapOfArea[i][j] + 1 {
			stack = append(stack, []int{i - 1, j})
		}
		if j + 1 < len(mapOfArea[0]) && mapOfArea[i][j + 1] == mapOfArea[i][j] + 1 {
			stack = append(stack, []int{i, j + 1})			
		}
		if j - 1 >= 0 && mapOfArea[i][j - 1] == mapOfArea[i][j] + 1 {
			stack = append(stack, []int{i, j - 1})			
		}
	}

	if shouldFindUnique {
		var uniqueReachablePositions = make(map[string]bool)
		for _, position := range reachablePositions {
			var key = strings.Join([]string{strconv.Itoa(position[0]), strconv.Itoa(position[1])}, ",")
			uniqueReachablePositions[key] = true
		}
		
		return len(uniqueReachablePositions)
	}

	return len(reachablePositions)
}

func solvePartOne(input [][]int) int {
	var sumOfScore = 0
	
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 0 {
				sumOfScore += findReachablePositions(i, j, input, true)
			}
		}
	}

	return sumOfScore
}

func solvePartTwo(input [][]int) int {
	var sumOfScore = 0
	
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 0 {
				sumOfScore += findReachablePositions(i, j, input, false)
			}
		}
	}

	return sumOfScore
}

func main() {
	var input [][]int

	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line = scanner.Text()
		var values []int
		for _, value := range line {
			var valueAsInt, _ = strconv.Atoi(string(value))
			values = append(values, valueAsInt)
		}
		input = append(input, values)
	}

	var partOneAnswer = solvePartOne(input)
	fmt.Println(partOneAnswer)

	var partTwoAnswer = solvePartTwo(input)
	fmt.Println(partTwoAnswer)
}
