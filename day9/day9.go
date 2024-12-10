package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	part1Solution := part1(filename)
	part2Solution := part2(filename)
	fmt.Println("Part 1:", part1Solution)
	fmt.Println("Part 2:", part2Solution)
}

const FREE int = -1

func part1(filename string) int {
	blocks := parseFile(filename)
	exploded := explodeFormat(blocks)

	for j := len(exploded) - 1; j >= 0; j-- {
		if exploded[j] == FREE {
			// ignore free
			continue
		}
		replaceIndex := findFirstFree(exploded)
		if replaceIndex >= j {
			continue
		}
		exploded[replaceIndex] = exploded[j]
		exploded[j] = FREE
	}

	return calcChecksum(exploded)
}

func part2(filename string) int {
	blocks := parseFile(filename)
	exploded := explodeFormat(blocks)
	for j := len(exploded) - 1; j >= 0; j-- {
		if exploded[j] == FREE {
			// ignore free
			continue
		}
		startFile, endFile := getSectorBounds(exploded, j)
		fileSize := endFile - startFile + 1

		freeSectorIndex := findFreeSector(exploded, fileSize)
		if freeSectorIndex < 0 || freeSectorIndex > j {
			continue
		}

		startFree, endFree := getSectorBounds(exploded, freeSectorIndex)
		freeSize := endFree - startFree + 1

		if fileSize > freeSize {
			continue
		}

		for freeIndex, fileIndex := startFree, startFile; freeIndex <= endFree && fileIndex <= endFile; freeIndex, fileIndex = freeIndex+1, fileIndex+1 {
			exploded[freeIndex] = exploded[fileIndex]
			exploded[fileIndex] = FREE
		}

	}

	return calcChecksum(exploded)
}

func explodeFormat(blocks []int) []int {
	exploded := make([]int, 0)

	blockId := 0
	for i, block := range blocks {
		if i%2 == 0 {
			// File
			for range block {
				exploded = append(exploded, blockId)
			}
			blockId += 1
		} else {
			// Free
			for range block {
				exploded = append(exploded, FREE)
			}
		}
	}
	return exploded
}

func getSectorBounds(data []int, index int) (int, int) {
	blockId := data[index]
	start := index
	end := index

	// Forward
	for i := index; i < len(data)-1; i++ {
		if data[i] != blockId {
			break
		}
		end = i
	}

	// Backward
	for i := index; i >= 0; i-- {
		if data[i] != blockId {
			break
		}
		start = i
	}

	return start, end
}

func calcChecksum(data []int) int {
	checksum := 0
	for i, num := range data {
		if num == FREE {
			continue
		}
		checksum += i * num
	}
	return checksum
}

func findFreeSector(blocks []int, size int) int {
	freeSize := 0
	for i, block := range blocks {
		if block == FREE {
			freeSize += 1
		} else {
			freeSize = 0
		}

		if freeSize == size {
			return i
		}

	}
	return -1
}

func findFirstFree(data []int) int {
	for i, block := range data {
		if block == FREE {
			return i
		}
	}
	return -1
}

func parseFile(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Failed to open file", filename)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	nums := make([]int, 0)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("Failed to read from file")
		}
		for _, c := range line {
			val, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatalln("Failed to parse number", err)
			}
			nums = append(nums, val)
		}
	}
	return nums
}
