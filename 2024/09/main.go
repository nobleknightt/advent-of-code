package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertDiskMapToDisk(diskMap string) []int {
	var disk []int

	var isFileLength = true
	var fileId = 0
	for _, length := range diskMap {
		var lengthAsInt, _ = strconv.Atoi(string(length))
		var value int
		if isFileLength {
			isFileLength = false
			value = fileId
			fileId++
		} else {
			isFileLength = true
			value = -1
		}
		for i := 0; i < lengthAsInt; i++ {
			disk = append(disk, value)
		}
	}

	return disk
}

func calculateCheckSum(disk []int) int {
	var checkSum = 0
	for i := 0; i < len(disk); i++ {
		if disk[i] != -1 {
			checkSum += i * disk[i]
		}
	}
	return checkSum
}

func solvePartOne(input string) int {

	var disk = convertDiskMapToDisk(input)

	var freeSpaceIndex, fileBlockIndex = 0, len(disk) - 1
	for freeSpaceIndex < fileBlockIndex {
		if disk[freeSpaceIndex] == -1 && disk[fileBlockIndex] != -1 {
			disk[freeSpaceIndex], disk[fileBlockIndex] = disk[fileBlockIndex], disk[freeSpaceIndex]
			freeSpaceIndex++
			fileBlockIndex--
		} else if disk[freeSpaceIndex] != -1 {
			freeSpaceIndex++
		} else if disk[fileBlockIndex] == -1 {
			fileBlockIndex--
		}
	}

	return calculateCheckSum(disk)
}

func solvePartTwo(input string) int {

	var disk = convertDiskMapToDisk(input)

	var freeSpaceIndex, fileBlockIndex = 0, len(disk) - 1
	for freeSpaceIndex < fileBlockIndex {
		var fileBlockLen = 0
		if disk[fileBlockIndex] != -1 {
			var fileId = disk[fileBlockIndex]
			for fileBlockIndex >= 0 && disk[fileBlockIndex] == fileId {
				fileBlockLen++
				fileBlockIndex--
			}
		} else {
			fileBlockIndex--
		}

		var tempFreeSpaceIndex = freeSpaceIndex
		for tempFreeSpaceIndex <= fileBlockIndex {
			var freeSpaceLen = 0
			if disk[tempFreeSpaceIndex] == -1 {
				for tempFreeSpaceIndex <= fileBlockIndex && disk[tempFreeSpaceIndex] == -1 {
					freeSpaceLen++
					tempFreeSpaceIndex++
				}

				if freeSpaceLen >= fileBlockLen {
					for i, j := tempFreeSpaceIndex-freeSpaceLen, fileBlockIndex+1; j < fileBlockIndex+1+fileBlockLen; i, j = i+1, j+1 {
						disk[i], disk[j] = disk[j], disk[i]
					}
					break
				}
			} else {
				tempFreeSpaceIndex++
			}
		}
	}

	return calculateCheckSum(disk)
}

func main() {
	var input string

	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	var partOneAnswer = solvePartOne(input)
	fmt.Println(partOneAnswer)

	var partTwoAnswer = solvePartTwo(input)
	fmt.Println(partTwoAnswer)
}
