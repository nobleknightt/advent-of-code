package main

import (
	"bufio"
	"fmt"
	"os"
)

func solvePartOne(schematics [][]string) int {
	var locks, keys [][]int
	for _, schematic := range schematics {
		var heights []int
		for i := 0; i < len(schematic[0]); i++ {
			heights = append(heights, -1)
		}
		for _, row := range schematic {
			for index, value := range row {
				if value == '#' {
					heights[index]++
				}
			}
		}
		if schematic[0][0] == '#' {
			locks = append(locks, heights)
		} else {
			keys = append(keys, heights)
		}
	}

	var numValidPairs = 0
	var height = len(schematics[0]) - 1

	for _, lock := range locks {
		for _, key := range keys {
			var isOverlapping = false
			for i := 0; i < len(lock); i++ {
				if lock[i] + key[i] >= height {
					isOverlapping = true
					break
				}
			}
			if !isOverlapping {
				numValidPairs++
			}
		}
	}

	return numValidPairs
}

func solvePartTwo(schematics [][]string) int {
	return len(schematics)
}

func main() {
	var schematics [][]string
	var schematic []string

	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line = scanner.Text()
		if len(line) > 0 {
			schematic = append(schematic, line)
		} else {
			schematics = append(schematics, schematic)
			schematic = []string{}
		}
	}
	schematics = append(schematics, schematic)

	if err := scanner.Err(); err != nil {
		return
	}

	var partOneAnswer = solvePartOne(schematics)
	fmt.Println(partOneAnswer)

	var partTwoAnswer = solvePartTwo(schematics)
	fmt.Println(partTwoAnswer)
}
