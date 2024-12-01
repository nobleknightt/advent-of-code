package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		return
	}
	defer file.Close()

	var left []int
	var count map[int]int
	count = make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		var value1, value2 int
		fmt.Sscanf(values[0], "%d", &value1)
		fmt.Sscanf(values[1], "%d", &value2)

		left = append(left, value1)
		count[value2]++
	}

	if err := scanner.Err(); err != nil {
		return
	}

	var similarityScore = 0
	for _, value := range left {
		similarityScore += value * count[value]
	}

	fmt.Println(similarityScore)
}
