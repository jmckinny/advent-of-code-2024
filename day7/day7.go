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

func part1(filename string) int {
	equations := parseFile(filename)

	total := 0
	for _, equation := range equations {
		if equation.trueViaAddMult() {
			total += equation.testValue
		}
	}

	return total
}

func part2(filename string) int {
	equations := parseFile(filename)

	total := 0
	for _, equation := range equations {
		if equation.trueViaAddMultConcat() {
			total += equation.testValue
		}
	}

	return total
}

type Operator int

const (
	ADD  Operator = 0
	MULT Operator = 1
)

type Equation struct {
	testValue int
	numbers   []int
}

func (self Equation) trueViaAddMult() bool {
	return self.trueViaAddMultHelper(1, self.numbers[0])
}

func (self Equation) trueViaAddMultHelper(index int, total int) bool {
	if index == len(self.numbers) {
		return total == self.testValue
	}
	return self.trueViaAddMultHelper(index+1, total*self.numbers[index]) || self.trueViaAddMultHelper(index+1, total+self.numbers[index])
}

func (self Equation) trueViaAddMultConcat() bool {
	val := self.numbers[0]
	remaining := RemoveIndex(self.numbers, 0)
	return self.trueViaAddMultConcatHelper(remaining, val)
}

func (self Equation) trueViaAddMultConcatHelper(curNumbers []int, total int) bool {
	if len(curNumbers) == 0 {
		return total == self.testValue
	}
	val := curNumbers[0]
	remaining := RemoveIndex(curNumbers, 0)

	multTrue := self.trueViaAddMultConcatHelper(remaining, total*val)
	if multTrue {
		return true
	}

	addTrue := self.trueViaAddMultConcatHelper(remaining, total+val)
	if addTrue {
		return true
	}

	str1 := strconv.Itoa(total)
	str2 := strconv.Itoa(curNumbers[0])
	concatNum := str1 + str2
	newNum, _ := strconv.Atoi(concatNum)
	concatTrue := self.trueViaAddMultConcatHelper(remaining, newNum)
	return concatTrue
}

func RemoveIndex(s []int, index int) []int {
	result := make([]int, len(s)-1)
	copy(result, s[:index])
	copy(result[index:], s[index+1:])
	return result
}

func parseFile(filename string) []Equation {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file", filename)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	equations := make([]Equation, 0)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("Failed to read from file")
		}

		splitByColon := strings.Split(line, ":")
		testValueString := splitByColon[0]
		testValue, err := strconv.Atoi(testValueString)
		if err != nil {
			log.Fatal("Failed to parse testValue ", err)
		}

		numbers := make([]int, 0)
		lineRest := strings.TrimSpace(splitByColon[1])
		for _, numStr := range strings.Split(lineRest, " ") {
			num, numErr := strconv.Atoi(numStr)

			if numErr != nil {
				log.Fatal("Failed to parse num ", err)
			}

			numbers = append(numbers, num)
		}

		equations = append(equations, Equation{testValue, numbers})
	}

	return equations
}
