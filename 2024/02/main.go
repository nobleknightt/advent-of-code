package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isIncreasing(report []int) bool {
	for i := 1; i < len(report); i++ {
		if report[i]-report[i-1] <= 0 {
			return false
		}
	}
	return true
}

func isDecreasing(report []int) bool {
	for i := 1; i < len(report); i++ {
		if report[i]-report[i-1] >= 0 {
			return false
		}
	}
	return true
}

func isDifferenceValid(report []int) bool {
	for i := 1; i < len(report); i++ {
		var diff = int64(math.Abs(float64(report[i] - report[i-1])))
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

// slices are reference types. This means that when you pass a slice to a function,
// you are passing a reference to the underlying array, not a copy of the slice.
// As a result, any modifications to the slice's elements or the slice itself
// (like changing its contents) will affect the original slice in the calling code.
func removeElement(s []int, index int) []int {
	var t []int
	t = append(t, s[:index]...)
	t = append(t, s[index+1:]...)
	return t
}

func isSafeAfterRemoval(report []int) bool {
	for i := 0; i < len(report); i++ {
		var updatedReport = removeElement(report, i)

		if (isIncreasing(updatedReport) || isDecreasing(updatedReport)) && isDifferenceValid(updatedReport) {
			return true
		}
	}

	return false
}

func solvePartOne(input [][]int) int {
	var numSafeReports = 0

	for _, report := range input {
		var isReportSafe = (isIncreasing(report) || isDecreasing(report)) && isDifferenceValid(report)
		if isReportSafe {
			numSafeReports++
		}
	}

	return numSafeReports
}

func solvePartTwo(input [][]int) int {
	var numSafeReports = 0

	for _, report := range input {
		var isReportSafe = (isIncreasing(report) || isDecreasing(report)) && isDifferenceValid(report)
		if isReportSafe || isSafeAfterRemoval(report) {
			numSafeReports++
		}
	}

	return numSafeReports
}

func main() {
	var input [][]int

	var scanner = bufio.NewScanner(os.Stdin)
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
