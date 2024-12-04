package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	part1Solution := part1()
	part2Solution := part2()
	fmt.Println("Part 1:", part1Solution)
	fmt.Println("Part 2:", part2Solution)
}

func part1() int {
	data := parseInput("input.txt", 1)
	total := 0
	for _, mul := range data {
		total += mul.compute()
	}
	return total
}

func part2() int {
	data := parseInput("input.txt", 2)
	total := 0
	for _, mul := range data {
		total += mul.compute()
	}
	return total
}

type Mult struct {
	x int
	y int
}

func (self Mult) compute() int {
	return self.x * self.y
}

func parseInput(filename string, partNum int) []Mult {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file", filename)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	content, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal("Failed to read input file")
	}
	switch partNum {
	case 1:
		return parseMults(string(content))
	case 2:
		return parseDoMults(string(content))
	default:
		log.Fatal("Not part 1 or part 2")
		return nil
	}
}

func parseMults(input string) []Mult {
	regexPattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := regexPattern.FindAllStringSubmatch(input, -1)
	mults := make([]Mult, 0)
	for _, match := range matches {
		x, errx := strconv.Atoi(match[1])
		y, erry := strconv.Atoi(match[2])
		if errx != nil || erry != nil {
			log.Fatal("Failed to parse ints")
		}
		mul := Mult{x, y}
		mults = append(mults, mul)
	}
	return mults
}

func parseDoMults(input string) []Mult {
	regexPattern := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	matches := regexPattern.FindAllStringSubmatch(input, -1)
	mults := make([]Mult, 0)
	valid := true
	for _, match := range matches {
		switch match[0] {

		case "do()":
			valid = true
		case "don't()":
			valid = false
		default:
			if valid {
				x, errx := strconv.Atoi(match[1])
				y, erry := strconv.Atoi(match[2])
				if errx != nil || erry != nil {
					log.Fatal("Failed to parse ints")
				}
				mul := Mult{x, y}
				mults = append(mults, mul)
			}
		}
	}
	return mults
}
