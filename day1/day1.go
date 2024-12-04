package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1Solution := part1()
	part2Solution := part2()
	fmt.Println("Part 1:", part1Solution)
	fmt.Println("Part 2:", part2Solution)
}

func loadNumbers(filename string) ([]int, []int) {
	lst1 := make([]int, 0)
	lst2 := make([]int, 0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file", filename)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("Failed to read from file")
		}
		split := strings.Fields(line)
		num1, _ := strconv.Atoi(split[0])
		num2, _ := strconv.Atoi(split[1])
		lst1 = append(lst1, num1)
		lst2 = append(lst2, num2)
	}

	return lst1, lst2
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func part1() int {
	lst1, lst2 := loadNumbers("input.txt")
	sort.Ints(lst1)
	sort.Ints(lst2)

	total := 0
	for i := range len(lst1) {
		diff := lst1[i] - lst2[i]
		total += abs(diff)
	}
	return total
}

func part2() int {
	lst1, lst2 := loadNumbers("input.txt")
	simularityScore := 0
	freqTable := createFrequencyTable(lst2)
	for i := range len(lst1) {
		simularityScore += lst1[i] * freqTable[lst1[i]]
	}
	return simularityScore
}

func createFrequencyTable(slice []int) map[int]int {
	table := make(map[int]int)
	for i := range len(slice) {
		itemCount := count(slice, slice[i])
		table[slice[i]] = itemCount
	}
	return table
}

func count(slice []int, item int) int {
	total := 0
	for i := range len(slice) {
		if slice[i] == item {
			total += 1
		}
	}
	return total
}
