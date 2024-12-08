package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type Location struct {
	x, y int
}

func isAntenna(r rune) bool {
	return unicode.IsLower(r) || unicode.IsUpper(r) || unicode.IsDigit(r)
}

func findAntinodesWithinBound(a, b Location, findAll bool, m, n int) []Location {
	var dx, dy = b.x - a.x, b.y - a.y

	var antinodes []Location
	var antinode = Location{a.x - dx, a.y - dy}

	for (0 <= antinode.x && antinode.x < m) && (0 <= antinode.y && antinode.y < n) {
		antinodes = append(antinodes, antinode)
		if !findAll {
			break
		}
		antinode = Location{antinode.x - dx, antinode.y - dy}
	}

	antinode = Location{b.x + dx, b.y + dy}
	for (0 <= antinode.x && antinode.x < m) && (0 <= antinode.y && antinode.y < n) {
		antinodes = append(antinodes, antinode)
		if !findAll {
			break
		}
		antinode = Location{antinode.x + dx, antinode.y + dy}
	}

	return antinodes
}

func solvePartOne(input []string) int {
	var antennaLocations = make(map[byte][]Location)

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			var char = input[i][j]
			if isAntenna(rune(char)) {
				antennaLocations[char] = append(antennaLocations[char], Location{i, j})
			}
		}
	}

	var antinodeLocations = make(map[Location]bool)

	for _, locations := range antennaLocations {
		for i := 0; i < len(locations); i++ {
			for j := i + 1; j < len(locations); j++ {
				var antinodes = findAntinodesWithinBound(locations[i], locations[j], false, len(input), len(input[0]))
				for _, antinode := range antinodes {
					antinodeLocations[antinode] = true
				}
			}
		}
	}

	return len(antinodeLocations)
}

func solvePartTwo(input []string) int {
	var antennaLocations = make(map[byte][]Location)

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			var char = input[i][j]
			if isAntenna(rune(char)) {
				antennaLocations[char] = append(antennaLocations[char], Location{i, j})
			}
		}
	}

	var antinodeLocations = make(map[Location]bool)

	for _, locations := range antennaLocations {
		for i := 0; i < len(locations); i++ {
			for j := i + 1; j < len(locations); j++ {
				antinodeLocations[locations[i]] = true
				antinodeLocations[locations[j]] = true

				var antinodes = findAntinodesWithinBound(locations[i], locations[j], true, len(input), len(input[0]))
				for _, antinode := range antinodes {
					antinodeLocations[antinode] = true
				}
			}
		}
	}

	return len(antinodeLocations)
}

func main() {
	var input []string

	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line = scanner.Text()
		input = append(input, line)
	}

	var partOneAnswer = solvePartOne(input)
	fmt.Println(partOneAnswer)

	var partTwoAnswer = solvePartTwo(input)
	fmt.Println(partTwoAnswer)
}
