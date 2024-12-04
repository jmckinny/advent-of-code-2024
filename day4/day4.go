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
	part1_solution := part1(data)
	part2_solution := part2(data)
	fmt.Println("Part 1:", part1_solution)
	fmt.Println("Part 2:", part2_solution)
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
			if ValidX_MAS(input, i, j) {
				total += 1
			}
		}
	}

	return total
}

func ValidX_MAS(input []string, row int, col int) bool {
	if get_rune_at(input, row, col) != 'A' {
		return false
	}

	offsets := []struct {
		row_dir, col_dir int
	}{
		{1, -1},  // Down + Left
		{1, 1},   // Down + Right
		{-1, 1},  // Up + Right
		{-1, -1}, // Up + Left
	}
	m_count := 0
	s_count := 0
	for _, dir := range offsets {
		if get_rune_at(input, row+dir.row_dir, col+dir.col_dir) == 'S' {
			s_count += 1
		}
		if get_rune_at(input, row+dir.row_dir, col+dir.col_dir) == 'M' {
			m_count += 1
		}
	}
	// Prevent:
	// S . M
	// . A .
	// M . S
	no_bad_arrange := get_rune_at(input, row+1, col+1) != get_rune_at(input, row-1, col-1)

	return m_count == 2 && s_count == 2 && no_bad_arrange
}

func CountValidXMAS(input []string, row int, col int) int {
	count := 0

	offsets := []struct {
		row_dir, col_dir int
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
		if get_rune_at(input, row+dir.row_dir, col+dir.col_dir) == 'M' &&
			get_rune_at(input, row+2*dir.row_dir, col+2*dir.col_dir) == 'A' &&
			get_rune_at(input, row+3*dir.row_dir, col+3*dir.col_dir) == 'S' {
			count++
		}
	}

	return count
}

func get_rune_at(input []string, row int, col int) rune {
	const INVALID = '.'

	height_valid := (row >= 0 && row < len(input))
	width_valid := (col >= 0 && col < len(input[0]))

	if !height_valid || !width_valid {
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
