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
	part1_solution := part1()
	part2_solution := part2()
	fmt.Println("Part 1:", part1_solution)
	fmt.Println("Part 2:", part2_solution)
}

func load_numbers(filename string) ([]int, []int) {
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
	lst1, lst2 := load_numbers("input.txt")
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
	lst1, lst2 := load_numbers("input.txt")
	simularity_score := 0
	freq_table := create_frequency_table(lst2)
	for i := range len(lst1) {
		simularity_score += lst1[i] * freq_table[lst1[i]]
	}
	return simularity_score
}

func create_frequency_table(slice []int) map[int]int {
	table := make(map[int]int)
	for i := range len(slice) {
		item_count := count(slice, slice[i])
		table[slice[i]] = item_count
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
