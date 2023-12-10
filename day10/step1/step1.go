package step1

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

	length := 2
	xy, dir := getFirstAfterStart(startXY, grid)
	var v rune
	for v != START_CHAR {
		v, xy, dir = getNextPipeCoords(xy, dir, grid)
		length++
	}

	return length / 2
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
