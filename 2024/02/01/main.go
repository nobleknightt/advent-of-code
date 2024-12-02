package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isSafe(prev, curr int) bool {
	if (prev < 0 && curr > 0) || (prev > 0 && curr < 0) {
		return false
	}

	var currAbs = int64(math.Abs(float64(curr)))
	if currAbs < 1 || currAbs > 3 {
		return false
	}

	return true
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

		var diff = report[0] - report[1]
		var diffAbs = int64(math.Abs(float64(diff)))
		if diffAbs < 1 || diffAbs > 3 {
			continue
		}

		var isReportSafe = true
		for i := 1; i < len(report); i++ {
			var currDiff = report[i-1] - report[i]
			if !isSafe(diff, currDiff) {
				isReportSafe = false
				break
			}
		}

		if isReportSafe {
			numSafeReports++
		}
	}

	if err := scanner.Err(); err != nil {
		return
	}

	fmt.Println(numSafeReports)
}
