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
	p1Answer := make(chan int)
	p2Answer := make(chan int)

	go part1(p1Answer)
	go part2(p2Answer)

	part1Solution := <-p1Answer
	part2Solution := <-p2Answer

	fmt.Println("Part 1:", part1Solution)
	fmt.Println("Part 2:", part2Solution)
}

func part1(answer chan<- int) {
	report := parseLevels("input.txt")
	totalSafe := 0
	for _, levels := range report {
		if reportIsSafe(levels) {
			totalSafe += 1
		}
	}
	answer <- totalSafe
}

func part2(answer chan<- int) {
	reports := parseLevels("input.txt")
	totalSafe := 0
	for _, levels := range reports {
		if reportIsSafeTolerant(levels) {
			totalSafe += 1
		}
	}

	answer <- totalSafe
}

func parseLevels(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file", filename)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	levels := make([][]int, 0)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("Failed to read from file")
		}

		level := make([]int, 0)
		split := strings.Fields(line)
		for i := range len(split) {
			num, err := strconv.Atoi(split[i])
			if err != nil {
				log.Fatal("Failed to parse number", err)
			}
			level = append(level, num)
		}
		levels = append(levels, level)
	}
	return levels
}

func reportIsSafe(report []int) bool {
	decreasing := report[0] > report[1]
	for i := 1; i < len(report); i++ {
		prev := report[i-1]
		cur := report[i]

		// all increase or all decrease
		if decreasing && cur > prev {
			return false
		} else if !decreasing && cur < prev {
			return false
		}
		// Gap must be at least 1 and at most 3
		gap := abs(cur - prev)
		if gap > 3 || gap < 1 {
			return false
		}

	}
	return true
}

func reportIsSafeTolerant(report []int) bool {
	if reportIsSafe(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		newLevels := RemoveIndex(report, i)
		if reportIsSafe(newLevels) {
			return true
		}
	}

	return false
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func RemoveIndex(s []int, index int) []int {
	result := make([]int, len(s)-1)
	copy(result, s[:index])
	copy(result[index:], s[index+1:])
	return result
}
