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

func removeElement(s []int, index int) []int {
	var t []int
	t = append(t, s[:index]...)
	t = append(t, s[index+1:]...)
	return t
}

// slices are reference types. This means that when you pass a slice to a function,
// you are passing a reference to the underlying array, not a copy of the slice.
// As a result, any modifications to the slice's elements or the slice itself
// (like changing its contents) will affect the original slice in the calling code.
func isSafeAfterRemoval(report []int) bool {
	for i := 0; i < len(report); i++ {
		var updatedReport = removeElement(report, i)

		if len(updatedReport) <= 1 {
			return true
		}

		if (isIncreasing(updatedReport) || isDecreasing(updatedReport)) && isDifferenceValid(updatedReport) {
			return true
		}
	}

	return false
}

func main() {
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

	var numSafeReports = 0

	for scanner.Scan() {
		var line = scanner.Text()
		var values = strings.Fields(line)

		var report []int
		for _, value := range values {
			var valueAsInt, _ = strconv.Atoi(value)
			report = append(report, valueAsInt)
		}

		var isSafe = (isIncreasing(report) || isDecreasing(report)) && isDifferenceValid(report)
		if isSafe || isSafeAfterRemoval(report) {
			numSafeReports++
		}
	}

	if err := scanner.Err(); err != nil {
		return
	}

	fmt.Println(numSafeReports)
}
