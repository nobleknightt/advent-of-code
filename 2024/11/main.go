package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solvePartOne(input []int) int {
	var numBlinks = 25
	for i := 0; i < numBlinks; i++ {
		var stonesAfterBlink []int
		for _, stone := range input {
			if stone == 0 {
				stonesAfterBlink = append(stonesAfterBlink, 1)
			} else {
				var stoneAsString = strconv.Itoa(stone)
				if len(stoneAsString)%2 == 0 {
					var left, _ = strconv.Atoi(stoneAsString[:len(stoneAsString)/2])
					var right, _ = strconv.Atoi(stoneAsString[len(stoneAsString)/2:])
					stonesAfterBlink = append(stonesAfterBlink, left)
					stonesAfterBlink = append(stonesAfterBlink, right)
				} else {
					stonesAfterBlink = append(stonesAfterBlink, stone*2024)
				}
			}
		}
		input = stonesAfterBlink
	}

	return len(input)
}

func solvePartTwo(input []int) int {
	var numBlinks = 75

	var stonesFrequency = make(map[int]int)
	for _, stone := range input {
		stonesFrequency[stone]++
	}

	for i := 0; i < numBlinks; i++ {
		var stonesFrequencyAfterBlink = make(map[int]int)
		for stone, frequency := range stonesFrequency {
			if stone == 0 {
				stonesFrequencyAfterBlink[1] += frequency
			} else {
				var stoneAsString = strconv.Itoa(stone)
				if len(stoneAsString)%2 == 0 {
					var left, _ = strconv.Atoi(stoneAsString[:len(stoneAsString)/2])
					var right, _ = strconv.Atoi(stoneAsString[len(stoneAsString)/2:])
					stonesFrequencyAfterBlink[left] += frequency
					stonesFrequencyAfterBlink[right] += frequency
					stonesFrequency[stone] = 0
				} else {
					stonesFrequencyAfterBlink[stone*2024] += frequency
				}
			}
			stonesFrequency[stone] = 0
		}
		stonesFrequency = stonesFrequencyAfterBlink
	}

	var numStonesAfterBlinks = 0
	for _, frequency := range stonesFrequency {
		numStonesAfterBlinks += frequency
	}

	return numStonesAfterBlinks
}

func main() {
	var input []int

	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line = scanner.Text()
		for _, value := range strings.Fields(line) {
			var valueAsInt, _ = strconv.Atoi(string(value))
			input = append(input, valueAsInt)
		}
	}

	var partOneAnswer = solvePartOne(input)
	fmt.Println(partOneAnswer)

	var partTwoAnswer = solvePartTwo(input)
	fmt.Println(partTwoAnswer)
}
