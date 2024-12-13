package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func solvePartOne(input [][2]int) int {
	var numTokens int

	for i := 0; i < len(input); i += 3 {
		var minTokens = math.Inf(1)

		var xA, yA = input[i][0], input[i][1]
		var xB, yB = input[i+1][0], input[i+1][1]
		var xPrize, yPrize = input[i+2][0], input[i+2][1]

		for a := 0; a < 100; a++ {
			for b := 0; b < 100; b++ {
				if xA*a+xB*b == xPrize && yA*a+yB*b == yPrize {
					minTokens = math.Min(minTokens, float64(a*3+b))
				}
			}
		}
		if minTokens < math.Inf(1) {
			numTokens += int(minTokens)
		}
	}

	return numTokens
}

func solvePartTwo(input [][2]int) int {
	var numTokens int

	for i := 0; i < len(input); i += 3 {
		var xA, yA = input[i][0], input[i][1]
		var xB, yB = input[i+1][0], input[i+1][1]
		var xPrize, yPrize = input[i+2][0] + 10000000000000, input[i+2][1] + 10000000000000

		// 

		// xA * a + xB * b = xPrize
		// yA * a + yB * b = yPrize
		//
		// a = (xPrize - xB * b) / xA
		// a = (yPrize - yB * b) / yA
		//
		// (xPrize - xB * b) / xA = (yPrize - yB * b) / yA
		// (xPrize - xB * b) * yA = (yPrize - yB * b) * xA
		//
		// xPrize * yA - xB * b * yA = yPrize * xA - yB * b * xA
		// xPrize * yA - yPrize * xA = xB * b * yA - yB * b * xA
		// xPrize * yA - yPrize * xA = b * (xB * yA - yB * xA)
		//
		// b = (xPrize * yA - yPrize * xA) / (xB * yA - yB * xA)

		var b = (xPrize*yA - yPrize*xA) / (xB*yA - yB*xA)
		var a = (xPrize - xB*b) / xA

		if xA*a+xB*b == xPrize && yA*a+yB*b == yPrize {
			numTokens += int(a*3 + b)
		}
	}

	return numTokens
}

func main() {
	var input [][2]int

	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line = scanner.Text()

		if len(line) == 0 {
			continue
		}
		var splits = strings.FieldsFunc(line, func(r rune) bool { return strings.ContainsRune(":,", r) })
		var x, _ = strconv.Atoi(strings.TrimFunc(splits[1], func(r rune) bool { return strings.ContainsRune("XY+= ", r) }))
		var y, _ = strconv.Atoi(strings.TrimFunc(splits[2], func(r rune) bool { return strings.ContainsRune("XY+= ", r) }))
		input = append(input, [2]int{x, y})
	}

	if err := scanner.Err(); err != nil {
		return
	}

	var partOneAnswer = solvePartOne(input)
	fmt.Println(partOneAnswer)

	var partTwoAnswer = solvePartTwo(input)
	fmt.Println(partTwoAnswer)
}
