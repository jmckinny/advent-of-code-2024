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
	rules, updates := ParseFile("input.txt")
	part1Solution := part1(rules, updates)
	part2Solution := part2(rules, updates)
	fmt.Println("Part 1:", part1Solution)
	fmt.Println("Part 2:", part2Solution)
}

func part1(rules []Rule, updates []Update) int {
	ruleMap := createRuleMap(rules)
	total := 0

	for _, update := range updates {
		if validUpdate(update, ruleMap) {
			total += getMiddleUpdate(update)
		}
	}

	return total
}

func part2(rules []Rule, updates []Update) int {
	ruleMap := createRuleMap(rules)
	total := 0
	for _, update := range updates {
		if !validUpdate(update, ruleMap) {
			sortUpdate(update, ruleMap)
			total += getMiddleUpdate(update)
		}
	}
	return total
}

type Rule struct {
	value int
	lower int
}

type Update []int

func sortUpdate(update Update, ruleMap map[int][]int) {
	sort.Slice(update, func(i, j int) bool {
		rules := ruleMap[update[i]]
		for _, rule := range rules {
			if rule == update[j] {
				return false
			}
		}
		return true
	})
}

func validUpdate(update Update, ruleMap map[int][]int) bool {
	for i, num := range update {
		rules := ruleMap[num]
		for _, ruleNum := range rules {
			index := findNumIndex(update, ruleNum)
			if index >= 0 && index < i {
				return false
			}
		}
	}
	return true
}

func findNumIndex(update Update, num int) int {
	for i, n := range update {
		if n == num {
			return i
		}
	}
	return -1
}

func getMiddleUpdate(lst Update) int {
	size := len(lst)
	return lst[size/2]
}

func createRuleMap(rules []Rule) map[int][]int {
	ruleMap := make(map[int][]int)
	for _, rule := range rules {
		data, exists := ruleMap[rule.value]
		if !exists {
			data = make([]int, 0)
		}
		data = append(data, rule.lower)
		ruleMap[rule.value] = data
	}

	return ruleMap
}

func ParseFile(filename string) ([]Rule, []Update) {
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

	fileString := string(content)

	sectionSplit := strings.Split(fileString, "\n\n")
	rulesUnparsed := sectionSplit[0]
	updatesUnparsed := sectionSplit[1]

	rules := make([]Rule, 0)
	updates := make([]Update, 0)

	for _, line := range strings.Split(rulesUnparsed, "\n") {
		halves := strings.Split(line, "|")
		lower, errLower := strconv.Atoi(halves[0])
		higher, errHigher := strconv.Atoi(halves[1])
		if errLower != nil || errHigher != nil {
			log.Fatal("Failed to parse rule numbers", err)
		}

		rule := Rule{lower, higher}
		rules = append(rules, rule)
	}

	for _, line := range strings.Split(strings.TrimSpace(updatesUnparsed), "\n") {
		update := make([]int, 0)
		for _, numberString := range strings.Split(line, ",") {
			num, err := strconv.Atoi(numberString)
			if err != nil {
				log.Fatal("Failed to parse update number: ", err)
			}
			update = append(update, num)
		}
		updates = append(updates, update)
	}

	return rules, updates
}
