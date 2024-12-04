package main

import (
	"bufio"
	"fmt"
	"os"
)

func solvePartOne(input []string) int {
	var xmasCount = 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if i+3 < len(input) && input[i][j] == 'X' && input[i+1][j] == 'M' && input[i+2][j] == 'A' && input[i+3][j] == 'S' {
				xmasCount++
			}
			if i-3 >= 0 && input[i][j] == 'X' && input[i-1][j] == 'M' && input[i-2][j] == 'A' && input[i-3][j] == 'S' {
				xmasCount++
			}
			if j+3 < len(input[i]) && input[i][j] == 'X' && input[i][j+1] == 'M' && input[i][j+2] == 'A' && input[i][j+3] == 'S' {
				xmasCount++
			}
			if j-3 >= 0 && input[i][j] == 'X' && input[i][j-1] == 'M' && input[i][j-2] == 'A' && input[i][j-3] == 'S' {
				xmasCount++
			}
			if i+3 < len(input) && j+3 < len(input[i]) && input[i][j] == 'X' && input[i+1][j+1] == 'M' && input[i+2][j+2] == 'A' && input[i+3][j+3] == 'S' {
				xmasCount++
			}
			if i-3 >= 0 && j-3 >= 0 && input[i][j] == 'X' && input[i-1][j-1] == 'M' && input[i-2][j-2] == 'A' && input[i-3][j-3] == 'S' {
				xmasCount++
			}
			if i+3 < len(input) && j-3 >= 0 && input[i][j] == 'X' && input[i+1][j-1] == 'M' && input[i+2][j-2] == 'A' && input[i+3][j-3] == 'S' {
				xmasCount++
			}
			if i-3 >= 0 && j+3 < len(input[i]) && input[i][j] == 'X' && input[i-1][j+1] == 'M' && input[i-2][j+2] == 'A' && input[i-3][j+3] == 'S' {
				xmasCount++
			}
		}
	}

	return xmasCount
}

func solvePartTwo(input []string) int {
	var xmasCount = 0

	for i := 1; i < len(input) - 1; i++ {
		for j := 1; j < len(input[i]) - 1; j++ {
			var masCount = 0
			if input[i - 1][j - 1] == 'M' && input[i][j] == 'A' && input[i + 1][j + 1] == 'S' {
				masCount++
			}
			if input[i - 1][j - 1] == 'S' && input[i][j] == 'A' && input[i + 1][j + 1] == 'M' {
				masCount++
			}
			if input[i + 1][j - 1] == 'S' && input[i][j] == 'A' && input[i - 1][j + 1] == 'M' {
				masCount++
			}
			if input[i + 1][j - 1] == 'M' && input[i][j] == 'A' && input[i - 1][j + 1] == 'S' {
				masCount++
			}

			if masCount == 2 {
				xmasCount++
			}
		}
	}

	return xmasCount
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
