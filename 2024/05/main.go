package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkIfInRightOrder(numbers []int, order map[string]bool, shouldSort bool) bool {
	var isInRightOrder = true
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			var wrongOrder = strconv.Itoa(numbers[j+1]) + "|" + strconv.Itoa(numbers[j])
			if order[wrongOrder] {
				isInRightOrder = false
				if shouldSort {
					numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				}
			}
		}
	}

	return isInRightOrder
}

func solvePartOne(pages [][]int, order map[string]bool) int {
	var middlePagesSum = 0

	for _, page := range pages {
		var isPageInRightOrder = checkIfInRightOrder(page, order, false)
		if isPageInRightOrder {
			middlePagesSum += page[len(page)/2]
		}
	}

	return middlePagesSum
}

func solvePartTwo(pages [][]int, order map[string]bool) int {
	var middlePagesSum = 0

	for _, page := range pages {
		var isPageInRightOrder = checkIfInRightOrder(page, order, true)
		if !isPageInRightOrder {
			middlePagesSum += page[len(page)/2]
		}
	}

	return middlePagesSum
}

func main() {
	var order = make(map[string]bool)
	var pages [][]int

	var scanner = bufio.NewScanner(os.Stdin)
	var isOrderInput = true

	for scanner.Scan() {
		var line = scanner.Text()
		if len(line) == 0 {
			isOrderInput = false
		} else {
			if isOrderInput {
				order[line] = true
			} else {
				var values []int
				for _, value := range strings.Split(line, ",") {
					var valueAsInt, _ = strconv.Atoi(value)
					values = append(values, valueAsInt)
				}
				pages = append(pages, values)
			}
		}
	}

	var partOneAnswer = solvePartOne(pages, order)
	fmt.Println(partOneAnswer)

	var partTwoAnswer = solvePartTwo(pages, order)
	fmt.Println(partTwoAnswer)
}
