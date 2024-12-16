package main

import (
	"bufio"
	"fmt"
	"os"
)

func findPositionOfRobot(warehouse [][]byte) (int, int) {
	for i := 0; i < len(warehouse); i++ {
		for j := 0; j < len(warehouse[i]); j++ {
			if warehouse[i][j] == '@' {
				return i, j
			}
		}
	}
	return -1, -1
}

func moveRobot(xRobot, yRobot int, move byte, warehouse [][]byte) (int, int) {
	if move == '<' {
		for y := yRobot; y >= 0; y-- {
			if warehouse[xRobot][y] == '#' {
				return xRobot, yRobot
			}
			if warehouse[xRobot][y] == '.' {
				for ; y < yRobot; y++ {
					warehouse[xRobot][y] = warehouse[xRobot][y+1]
				}
				warehouse[xRobot][yRobot] = '.'
				return xRobot, yRobot - 1
			}
		}
	} else if move == '>' {
		for y := yRobot; y < len(warehouse[0]); y++ {
			if warehouse[xRobot][y] == '#' {
				return xRobot, yRobot
			}
			if warehouse[xRobot][y] == '.' {
				for ; y > yRobot; y-- {
					warehouse[xRobot][y] = warehouse[xRobot][y-1]
				}
				warehouse[xRobot][yRobot] = '.'
				return xRobot, yRobot + 1
			}
		}
	} else if move == '^' {
		for x := xRobot; x >= 0; x-- {
			if warehouse[x][yRobot] == '#' {
				return xRobot, yRobot
			}
			if warehouse[x][yRobot] == '.' {
				for ; x < xRobot; x++ {
					warehouse[x][yRobot] = warehouse[x+1][yRobot]
				}
				warehouse[xRobot][yRobot] = '.'
				return xRobot - 1, yRobot
			}
		}
	} else if move == 'v' {
		for x := xRobot; x < len(warehouse); x++ {
			if warehouse[x][yRobot] == '#' {
				return xRobot, yRobot
			}
			if warehouse[x][yRobot] == '.' {
				for ; x > xRobot; x-- {
					warehouse[x][yRobot] = warehouse[x-1][yRobot]
				}
				warehouse[xRobot][yRobot] = '.'
				return xRobot + 1, yRobot
			}
		}
	}
	return -1, -1
}

func plotWarehouse(warehouse [][]byte) {
	for i := 0; i < len(warehouse); i++ {
		for j := 0; j < len(warehouse[i]); j++ {
			fmt.Print(string(warehouse[i][j]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func solvePartOne(input []string) int {

	var warehouse [][]byte
	var moves []byte
	for i := 0; i < len(input); i++ {
		if input[i][0] == '#' {
			var row []byte
			for j := 0; j < len(input[i]); j++ {
				row = append(row, input[i][j])
			}
			warehouse = append(warehouse, row)
		} else {
			for j := 0; j < len(input[i]); j++ {
				moves = append(moves, input[i][j])
			}
		}
	}

	var xRobot, yRobot = findPositionOfRobot(warehouse)

	for _, move := range moves {
		xRobot, yRobot = moveRobot(xRobot, yRobot, move, warehouse)
	}

	plotWarehouse(warehouse)

	var sumOfGPSCoordinates int

	for i := 0; i < len(warehouse); i++ {
		for j := 0; j < len(warehouse[i]); j++ {
			if warehouse[i][j] == 'O' {
				var gpsCoordinate = 100*i + j
				sumOfGPSCoordinates += gpsCoordinate
			}
		}
	}

	return sumOfGPSCoordinates
}

func solvePartTwo(input []string) int {

	var warehouse [][]byte
	var moves []byte
	for i := 0; i < len(input); i++ {
		if input[i][0] == '#' {
			var row []byte
			for j := 0; j < len(input[i]); j++ {
				if input[i][j] == '#' {
					row = append(row, '#', '#')
				} else if input[i][j] == 'O' {
					row = append(row, '[', ']')
				} else if input[i][j] == '.' {
					row = append(row, '.', '.')
				} else if input[i][j] == '@' {
					row = append(row, '@', '.')
				}
			}
			warehouse = append(warehouse, row)
		} else {
			for j := 0; j < len(input[i]); j++ {
				moves = append(moves, input[i][j])
			}
		}
	}

	plotWarehouse(warehouse)

	return len(warehouse)
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
