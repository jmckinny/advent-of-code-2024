package main

import "testing"

func TestPart1(t *testing.T) {
	input_data := ParseFile("test_input.txt")
	part1_sol := part1(input_data)
	if part1_sol != 18 {
		t.Fatalf("Part 1 failed, expected %d got: %d", 18, part1_sol)
	}
}

func TestPart2(t *testing.T) {
	input_data := ParseFile("test_input.txt")
	part2_sol := part2(input_data)
	if part2_sol != 9 {
		t.Fatalf("Part 2 failed, expected %d got: %d", 9, part2_sol)
	}
}
