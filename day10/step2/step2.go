package step2

import (
	"strings"
)

const (
	START_CHAR = 'S'
	DIR_UP     = 'U'
	DIR_DOWN   = 'D'
	DIR_LEFT   = 'L'
	DIR_RIGHT  = 'R'
)

func Solve(input string) int {
	grid := [][]rune{}
	var startXY [2]int
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
		startIdx := strings.IndexRune(line, START_CHAR)
		if startIdx != -1 {
			startXY = [2]int{startIdx, len(grid) - 1}
		}
	}

	pipeCoords := make(map[int][]bool, len(grid))
	pipeCoords[startXY[1]] = make([]bool, len(grid[startXY[1]]))
	pipeCoords[startXY[1]][startXY[0]] = true

	xy, dir := getFirstAfterStart(startXY, grid)
	var v rune
	for v != START_CHAR {
		if _, ok := pipeCoords[xy[1]]; !ok {
			pipeCoords[xy[1]] = make([]bool, len(grid[xy[1]]))
		}
		pipeCoords[xy[1]][xy[0]] = true
		v, xy, dir = getNextPipeCoords(xy, dir, grid)
	}

	sum := 0
	for y := 0; y < len(grid); y++ {
		if _, ok := pipeCoords[y]; !ok {
			continue
		}

		line := ""
		for x, char := range grid[y] {
			if pipeCoords[y][x] {
				if char != '-' {
					line += string(char)
				}
			} else {
				line += "."
			}
		}

		line = strings.ReplaceAll(line, "L7", "|")
		line = strings.ReplaceAll(line, "FJ", "|")
		line = strings.ReplaceAll(line, "LJ", "")
		line = strings.ReplaceAll(line, "F7", "")

		isOdd := false
		for _, char := range line {
			if char == '|' {
				isOdd = !isOdd
			} else if char == '.' && isOdd {
				sum++
			}
		}
	}

	return sum
}

func getFirstAfterStart(startXY [2]int, grid [][]rune) ([2]int, rune) {
	if grid[startXY[1]][startXY[0]+1] == '-' || grid[startXY[1]][startXY[0]+1] == '7' {
		return [2]int{startXY[0] + 1, startXY[1]}, DIR_RIGHT
	}
	if grid[startXY[1]][startXY[0]-1] == '-' || grid[startXY[1]][startXY[0]-1] == 'F' {
		return [2]int{startXY[0] - 1, startXY[1]}, DIR_LEFT
	}
	if grid[startXY[1]+1][startXY[0]] == '|' || grid[startXY[1]+1][startXY[0]] == 'J' || grid[startXY[1]+1][startXY[0]] == 'L' {
		return [2]int{startXY[0], startXY[1] + 1}, DIR_DOWN
	}

	return [2]int{startXY[0], startXY[1] - 1}, DIR_UP
}

func getNextPipeCoords(xy [2]int, dir rune, grid [][]rune) (pipe rune, newXY [2]int, newDir rune) {
	pipe = grid[xy[1]][xy[0]]

	newDir = dir
	switch pipe {
	case '|':
		if dir == DIR_UP {
			newXY = [2]int{xy[0], xy[1] - 1}
		} else {
			newXY = [2]int{xy[0], xy[1] + 1}
		}
	case '-':
		if dir == DIR_LEFT {
			newXY = [2]int{xy[0] - 1, xy[1]}
		} else {
			newXY = [2]int{xy[0] + 1, xy[1]}
		}
	case 'J':
		if dir == DIR_DOWN {
			newXY = [2]int{xy[0] - 1, xy[1]}
			newDir = DIR_LEFT
		} else {
			newXY = [2]int{xy[0], xy[1] - 1}
			newDir = DIR_UP
		}
	case 'L':
		if dir == DIR_DOWN {
			newXY = [2]int{xy[0] + 1, xy[1]}
			newDir = DIR_RIGHT
		} else {
			newXY = [2]int{xy[0], xy[1] - 1}
			newDir = DIR_UP
		}
	case 'F':
		if dir == DIR_LEFT {
			newXY = [2]int{xy[0], xy[1] + 1}
			newDir = DIR_DOWN
		} else {
			newXY = [2]int{xy[0] + 1, xy[1]}
			newDir = DIR_RIGHT
		}
	case '7':
		if dir == DIR_RIGHT {
			newXY = [2]int{xy[0], xy[1] + 1}
			newDir = DIR_DOWN
		} else {
			newXY = [2]int{xy[0] - 1, xy[1]}
			newDir = DIR_LEFT
		}
	}

	pipe = grid[newXY[1]][newXY[0]]

	return
}
