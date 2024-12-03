package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solvePartOne(input [][]int) int {
	var left, right []int

	for _, line := range input {
		left = append(left, line[0])
		right = append(right, line[1])
	}

	sort.Ints(left)
	sort.Ints(right)

	var totalDistance = 0
	for i := 0; i < len(left); i++ {
		totalDistance += int(math.Abs(float64(left[i] - right[i])))
	}

	return totalDistance
}

func solvePartTwo(input [][]int) int {
	var left []int

	var count map[int]int
	count = make(map[int]int)

	for _, line := range input {
		left = append(left, line[0])
		count[line[1]]++
	}

	var similarityScore = 0
	for _, value := range left {
		similarityScore += value * count[value]
	}

	return similarityScore
}

func main() {
	var input [][]int

	var scanner = bufio.NewScanner(os.Stdin)
	// By default, bufio.Scanner uses bufio.ScanLines as the split function,
	// which means it reads the input line by line. Each time you call scanner.Scan(),
	// it reads a single line of input (up to a newline character \n), and you can
	// access the content of that line using scanner.Text().
	//
	// The Scan() method returns true if there is another token (i.e., a line) to
	// read, and it returns false if there are no more tokens to read, which could
	// happen either because the end of the input has been reached or an error occurred.
	//
	// After Scan() returns false, you can check for errors by calling scanner.Err().
	// If scanner.Err() returns nil, it means the input ended normally (EOF),
	// otherwise, it indicates an error occurred during scanning.

	for scanner.Scan() {
		var line = scanner.Text()
		var values []int
		for _, value := range strings.Fields(line) {
			var valueAsInt, _ = strconv.Atoi(value)
			values = append(values, valueAsInt)
		}
		input = append(input, values)
	}

	if err := scanner.Err(); err != nil {
		return
	}

	var partOneAnswer = solvePartOne(input)
	fmt.Println(partOneAnswer)

	var partTwoAnswer = solvePartTwo(input)
	fmt.Println(partTwoAnswer)
}
