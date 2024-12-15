package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func countRobotsPerQuadrant(robots [][4]int, width, height int) [4]int {
	var robotsPerQuadrant = [4]int{0, 0, 0, 0}

	for i := 0; i < len(robots); i++ {
		var xP, yP = robots[i][0], robots[i][1]
		if (0 <= xP && xP < width/2) && (0 <= yP && yP < height/2) {
			robotsPerQuadrant[0]++
		} else if (0 <= xP && xP < width/2) && (height/2 < yP && yP < height) {
			robotsPerQuadrant[1]++
		} else if (width/2 < xP && xP < width) && (0 <= yP && yP < height/2) {
			robotsPerQuadrant[2]++
		} else if (width/2 < xP && xP < width) && (height/2 < yP && yP < height) {
			robotsPerQuadrant[3]++
		}
	}

	return robotsPerQuadrant
}

func allAtDifferentPositions(robots [][4]int) bool {
	var occupied = make(map[string]bool)
	for i := 0; i < len(robots); i++ {
		var xP, yP = robots[i][0], robots[i][1]
		var position = strconv.Itoa(xP) + "," + strconv.Itoa(yP)
		if occupied[position] {
			return false
		}
		occupied[position] = true
	}
	return true
}

func plotRobots(robots [][4]int, width, height int) {
	var plot = make([][]bool, height)
	for row := range plot {
		plot[row] = make([]bool, width)
	}

	for i := 0; i < len(robots); i++ {
		var xP, yP = robots[i][0], robots[i][1]
		plot[yP][xP] = true
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if plot[i][j] {
				fmt.Print("\033[32m*\033[0m")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func simulateRobots(robots [][4]int, numSeconds, width, height int, shouldCheckForEasterEgg bool) int {
	for second := 0; second < numSeconds; second++ {
		if shouldCheckForEasterEgg && allAtDifferentPositions(robots) {
			plotRobots(robots, width, height)
			return second
		}

		for j := 0; j < len(robots); j++ {
			robots[j][0] = ((robots[j][0]+robots[j][2])%width + width) % width
			robots[j][1] = ((robots[j][1]+robots[j][3])%height + height) % height
		}
	}

	return 0
}

func solvePartOne(input [][4]int) int {
	var robots = make([][4]int, len(input))
	copy(robots, input)

	simulateRobots(robots, 100, 101, 103, false)
	var robotsPerQuadrant = countRobotsPerQuadrant(robots, 101, 103)

	return robotsPerQuadrant[0] * robotsPerQuadrant[1] * robotsPerQuadrant[2] * robotsPerQuadrant[3]
}

func solvePartTwo(input [][4]int) int {
	var robots = make([][4]int, len(input))
	copy(robots, input)

	return simulateRobots(robots, math.MaxInt, 101, 103, true)
}

func main() {
	var input [][4]int

	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line = scanner.Text()

		var splits = strings.FieldsFunc(line, func(r rune) bool { return strings.ContainsRune("=, ", r) })

		var xP, _ = strconv.Atoi(splits[1])
		var yP, _ = strconv.Atoi(splits[2])
		var xV, _ = strconv.Atoi(splits[4])
		var yV, _ = strconv.Atoi(splits[5])
		input = append(input, [4]int{xP, yP, xV, yV})
	}

	if err := scanner.Err(); err != nil {
		return
	}

	var partOneAnswer = solvePartOne(input)
	fmt.Println(partOneAnswer)

	var partTwoAnswer = solvePartTwo(input)
	fmt.Println(partTwoAnswer)
}
