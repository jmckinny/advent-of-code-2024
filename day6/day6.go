package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	filename := "input.txt"
	part1Solution := part1(filename)
	part2Solution := part2(filename)
	fmt.Println("Part 1:", part1Solution)
	fmt.Println("Part 2:", part2Solution)
}

type (
	Direction int
	MapTile   int
	Map       [][]MapTile
)

const (
	OPEN       MapTile = 0
	OBSTRUCTED MapTile = 1
	GUARD      MapTile = 2
	SEEN       MapTile = 3
)

const (
	UP    Direction = 0
	RIGHT Direction = 1
	DOWN  Direction = 2
	LEFT  Direction = 3
)

func part1(filename string) int {
	mapGrid := parseFile(filename)
	direction := UP
	row, col := mapGrid.getGaurdPos()

	for row != -1 || col != -1 {

		nextRow, nextCol, nextDir := mapGrid.getGaurdMove(row, col, direction)
		if nextRow == -1 && nextCol == -1 && nextDir == -1 {
			mapGrid[row][col] = SEEN
			break
		}
		direction = nextDir
		mapGrid[row][col] = SEEN
		mapGrid[nextRow][nextCol] = GUARD
		row = nextRow
		col = nextCol
	}
	return mapGrid.countSeen()
}

func part2(filename string) int {
	mapGrid := parseFile(filename)
	positions := 0
	jobs := make([]chan bool, 0)
	for row := range mapGrid {
		for col := range mapGrid[row] {
			if mapGrid[row][col] != OPEN {
				continue
			}
			mapCopy := mapGrid.clone()
			mapCopy[row][col] = OBSTRUCTED
			job := make(chan bool)
			jobs = append(jobs, job)
			go isLikelyInfinite(mapCopy, job)
		}
	}

	for _, job := range jobs {
		jobResult := <-job
		if jobResult {
			positions += 1
		}
	}

	return positions
}

func isLikelyInfinite(mapGrid Map, resultChannel chan bool) {
	maxSteps := len(mapGrid)*len(mapGrid[0]) - len(mapGrid)
	step := 0
	direction := UP
	row, col := mapGrid.getGaurdPos()

	for {
		if step > maxSteps {
			resultChannel <- true
			return
		}

		if row == -1 && col == -1 {
			break
		}

		nextRow, nextCol, nextDir := mapGrid.getGaurdMove(row, col, direction)
		if nextRow == -1 && nextCol == -1 && nextDir == -1 {
			break
		}
		direction = nextDir
		row = nextRow
		col = nextCol
		step += 1
	}
	resultChannel <- false
}

func (self Map) clone() Map {
	duplicate := make([][]MapTile, len(self))
	for i := range self {
		duplicate[i] = make([]MapTile, len(self[i]))
		copy(duplicate[i], self[i])
	}
	return duplicate
}

func (self Map) inBounds(row int, col int) bool {
	heightValid := (row >= 0 && row < len(self))
	widthValid := (col >= 0 && col < len(self[0]))

	if !heightValid || !widthValid {
		return false
	}
	return true
}

func (self Map) isObstruction(row int, col int) bool {
	heightValid := (row >= 0 && row < len(self))
	widthValid := (col >= 0 && col < len(self[0]))

	if !heightValid || !widthValid {
		return false
	}

	return self[row][col] == OBSTRUCTED
}

func (self Map) getGaurdMove(curRow int, curCol int, currentDir Direction) (int, int, Direction) {
	offsets := map[Direction]struct {
		rowDir, colDir int
	}{
		UP:    {-1, 0},
		DOWN:  {1, 0},
		LEFT:  {0, -1},
		RIGHT: {0, 1},
	}

	offset := offsets[currentDir]
	newRow := curRow + offset.rowDir
	newCol := curCol + offset.colDir
	if self.inBounds(newRow, newCol) {
		if self.isObstruction(newRow, newCol) {
			return curRow, curCol, (currentDir + 1) % 4
		} else {
			return newRow, newCol, currentDir
		}
	} else {
		return -1, -1, -1
	}
}

func (self Map) getGaurdPos() (int, int) {
	for i, row := range self {
		for j, tile := range row {
			if tile == GUARD {
				return i, j
			}
		}
	}
	return -1, -1
}

func (self Map) countSeen() int {
	count := 0
	for _, row := range self {
		for _, tile := range row {
			if tile == SEEN {
				count += 1
			}
		}
	}
	return count
}

func (self Map) printMap() {
	for _, row := range self {
		for _, tile := range row {
			switch tile {
			case OPEN:
				fmt.Print(".")
			case OBSTRUCTED:
				fmt.Print("#")
			case SEEN:
				fmt.Print("X")
			case GUARD:
				fmt.Print("G")
			}
		}
		fmt.Println()
	}
}

func parseFile(filename string) Map {
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
	inputData := string(content)

	mapLayout := make([][]MapTile, 0)
	for _, line := range strings.Split(strings.TrimSpace(inputData), "\n") {
		row := make([]MapTile, 0)
		for _, c := range line {
			switch c {
			case '.':
				row = append(row, OPEN)
			case '#':
				row = append(row, OBSTRUCTED)
			case '^':
				row = append(row, GUARD)
			default:
				log.Fatal("Unrecognized tile", c)
			}
		}
		mapLayout = append(mapLayout, row)
	}
	return mapLayout
}
