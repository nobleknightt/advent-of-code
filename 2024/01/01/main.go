package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		return
	}
	defer file.Close()

	var left []int
	var right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		var value1, value2 int
		fmt.Sscanf(values[0], "%d", &value1)
		fmt.Sscanf(values[1], "%d", &value2)

		left = append(left, value1)
		right = append(right, value2)
	}

	if err := scanner.Err(); err != nil {
		return
	}

	sort.Ints(left)
	sort.Ints(right)

	var totalDistance = 0
	for i := 0; i < len(left); i++ {
		totalDistance += int(math.Abs(float64(left[i] - right[i])))
	}

	fmt.Println(totalDistance)
}
