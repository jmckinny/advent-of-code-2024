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
	p1_answer := make(chan int)
	p2_answer := make(chan int)

	go part1(p1_answer)
	go part2(p2_answer)

	part1_solution := <-p1_answer
	part2_solution := <-p2_answer

	fmt.Println("Part 1:", part1_solution)
	fmt.Println("Part 2:", part2_solution)
}

func part1(answer chan<- int) {
	report := parse_levels("input.txt")
	total_safe := 0
	for _, levels := range report {
		if report_is_safe(levels) {
			total_safe += 1
		}
	}
	answer <- total_safe
}

func part2(answer chan<- int) {
	reports := parse_levels("input.txt")
	total_safe := 0
	for _, levels := range reports {
		if report_is_safe_tolerant(levels) {
			total_safe += 1
		}
	}

	answer <- total_safe
}

func parse_levels(filename string) [][]int {
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

func report_is_safe(report []int) bool {
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

func report_is_safe_tolerant(report []int) bool {
	if report_is_safe(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		new_levels := RemoveIndex(report, i)
		if report_is_safe(new_levels) {
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
