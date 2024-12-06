package main

import "testing"

func TestPart1(t *testing.T) {
	rules, updates := ParseFile("test_input.txt")
	part1Solution := 143
	part1Answer := part1(rules, updates)
	if part1Answer != part1Solution {
		t.Fatalf("Part 1 failed, expected %d got: %d", part1Solution, part1Answer)
	}
}

func TestPart2(t *testing.T) {
	rules, updates := ParseFile("test_input.txt")
	partSolution := 123
	partAnswer := part2(rules, updates)
	if partAnswer != partSolution {
		t.Fatalf("Part 2 failed, expected %d got: %d", partSolution, partAnswer)
	}
}
