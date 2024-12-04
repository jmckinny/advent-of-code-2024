package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	data := ParseFile("input.txt")
	part1Solution := part1(data)
	part2Solution := part2(data)
	fmt.Println("Part 1:", part1Solution)
	fmt.Println("Part 2:", part2Solution)
}

func part1(input []string) int {
	total := 0
	for i, row := range input {
		for j, c := range row {
			if c == 'X' {
				total += CountValidXMAS(input, i, j)
			}
		}
	}
	return total
}

func part2(input []string) int {
	total := 0
	for i, row := range input {
		for j := range len(row) {
			if ValidCrossMAS(input, i, j) {
				total += 1
			}
		}
	}

	return total
}

func ValidCrossMAS(input []string, row int, col int) bool {
	if getRuneAt(input, row, col) != 'A' {
		return false
	}

	offsets := []struct {
		rowDir, colDir int
	}{
		{1, -1},  // Down + Left
		{1, 1},   // Down + Right
		{-1, 1},  // Up + Right
		{-1, -1}, // Up + Left
	}
	mCount := 0
	sCount := 0
	for _, dir := range offsets {
		if getRuneAt(input, row+dir.rowDir, col+dir.colDir) == 'S' {
			sCount += 1
		}
		if getRuneAt(input, row+dir.rowDir, col+dir.colDir) == 'M' {
			mCount += 1
		}
	}
	// Prevent:
	// S . M
	// . A .
	// M . S
	noBadArrange := getRuneAt(input, row+1, col+1) != getRuneAt(input, row-1, col-1)

	return mCount == 2 && sCount == 2 && noBadArrange
}

func CountValidXMAS(input []string, row int, col int) int {
	count := 0

	offsets := []struct {
		rowDir, colDir int
	}{
		{-1, 0},  // Up
		{1, 0},   // Down
		{0, 1},   // Right
		{0, -1},  // Left
		{1, -1},  // Down + Left
		{1, 1},   // Down + Right
		{-1, 1},  // Up + Right
		{-1, -1}, // Up + Left
	}

	for _, dir := range offsets {
		if getRuneAt(input, row+dir.rowDir, col+dir.colDir) == 'M' &&
			getRuneAt(input, row+2*dir.rowDir, col+2*dir.colDir) == 'A' &&
			getRuneAt(input, row+3*dir.rowDir, col+3*dir.colDir) == 'S' {
			count++
		}
	}

	return count
}

func getRuneAt(input []string, row int, col int) rune {
	const INVALID = '.'

	heightValid := (row >= 0 && row < len(input))
	widthValid := (col >= 0 && col < len(input[0]))

	if !heightValid || !widthValid {
		return INVALID
	}

	return rune(input[row][col])
}

func ParseFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file", filename)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	lines := make([]string, 0)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("Failed to read from file")
		}
		lines = append(lines, line)
	}

	return lines
}
