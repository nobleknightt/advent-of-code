package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func isConnected(computerOne string, computers []string) bool {
	for _, computerTwo := range computers {
		if computerOne == computerTwo {
			return true
		}
	}
	return false
}

func isValidClique(clique []string, networkMap map[string][]string) bool {
	for _, computerOne := range clique {
		for _, computerTwo := range clique {
			if computerOne != computerTwo && !isConnected(computerOne, networkMap[computerTwo]) {
				return false
			}
		}
	}
	return true
}

func findMaximumClique(index int, clique []string, computers []string, networkMap map[string][]string) []string {
	if index == len(computers) {
		return clique
	}

	var maximamClique []string

	var newClique = make([]string, len(clique))
	copy(newClique, clique)

	newClique = append(newClique, computers[index])
	if isValidClique(newClique, networkMap) {
		maximamClique = findMaximumClique(index + 1, newClique, computers, networkMap)
	}

	var maximalClique = findMaximumClique(index + 1, clique, computers, networkMap)
	if len(maximalClique) > len(maximamClique) {
		maximamClique = maximalClique
	}

	return maximamClique
}

func solvePartOne(input []string) int {
	var networkMap = make(map[string][]string)
	for _, connection := range input {
		var computerOne = connection[:2]
		var computerTwo = connection[3:]
		networkMap[computerOne] = append(networkMap[computerOne], computerTwo)
		networkMap[computerTwo] = append(networkMap[computerTwo], computerOne)
	}

	var isCheckedGroup = make(map[string]bool)
	var numGroups = 0

	for computerOne := range networkMap {
		for _, computerTwo := range networkMap[computerOne] {
			for _, computerThree := range networkMap[computerTwo] {
				if isConnected(computerOne, networkMap[computerThree]) {
					if computerOne[0] == 't' || computerTwo[0] == 't' || computerThree[0] == 't' {
						var group = []string{computerOne, computerTwo, computerThree}
						sort.Strings(group)
						var groupKey = strings.Join(group, ",")
						if !isCheckedGroup[groupKey] {
							isCheckedGroup[groupKey] = true
							numGroups++
						}
					}
				}
			}
		}
	}

	return numGroups
}

func solvePartTwo(input []string) string {
	var networkMap = make(map[string][]string)
	for _, connection := range input {
		var computerOne = connection[:2]
		var computerTwo = connection[3:]
		networkMap[computerOne] = append(networkMap[computerOne], computerTwo)
		networkMap[computerTwo] = append(networkMap[computerTwo], computerOne)
	}

	// A clique in a graph is a subset of vertices such that every two distinct vertices
	// in the subset are adjacent (i.e., there's an edge between every pair).

	var computers []string
	for computer := range networkMap {
		computers = append(computers, computer)
	}

	var maximamClique = findMaximumClique(0, []string{}, computers, networkMap)
	sort.Strings(maximamClique)

	var password = strings.Join(maximamClique, ",")
	return password
}

func main() {
	var input []string

	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line = scanner.Text()
		input = append(input, line)
	}

	if err := scanner.Err(); err != nil {
		return
	}

	var partOneAnswer = solvePartOne(input)
	fmt.Println(partOneAnswer)

	var partTwoAnswer = solvePartTwo(input)
	fmt.Println(partTwoAnswer)
}
