package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func positionOfGuard(input []string) (int, int) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == '^' {
				return i, j
			}
		}
	}
	return -1, -1
}

func solvePartOne(input []string) int {
	var m, n = len(input), len(input[0])
	var x, y = positionOfGuard(input)

	var visitedPositions = make(map[string]bool)

	var directions = [4][2]int{{-1, 0}, {0, +1}, {+1, 0}, {0, -1}}
	var directionIndex = 0
	
	for (x >= 0 && x < m) && (y >= 0 && y < n) {
		var key = strings.Join([]string{strconv.Itoa(x),strconv.Itoa(y)}, ",")
		visitedPositions[key] = true

		if (x == 0 || x == m - 1) || (y == 0 || y == n - 1) {
			break
		}
				
		var xNext = x + directions[directionIndex][0]
		var yNext = y + directions[directionIndex][1]

		if (xNext >= 0 && xNext < m) && (yNext >= 0 && yNext < n) {
			if input[xNext][yNext] == '#' {
				directionIndex = (directionIndex + 1) % 4
				xNext = x + directions[directionIndex][0]
				yNext = y + directions[directionIndex][1]
			}
		}
		x, y = xNext, yNext
	} 

	return len(visitedPositions)
}

func solvePartTwo(input []string) int {
	return len(input)
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
