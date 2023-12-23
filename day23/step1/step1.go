package step1

import (
	"bytes"
	"slices"
)

var dirs = map[byte][2]int16{
	'^': {0, -1},
	'>': {1, 0},
	'v': {0, 1},
	'<': {-1, 0},
}

func Solve(input string) int {
	tiles := bytes.Fields([]byte(input))
	passed := make([][]int16, len(tiles))
	trail := make([]bool, len(tiles)*len(tiles[0]))
	for i := range passed {
		passed[i] = make([]int16, len(tiles[0]))
	}
	point := [2]int16{1, 0}
	end := [2]int16{int16(len(tiles[0]) - 2), int16(len(tiles) - 1)}

	return int(getTileCount(point, end, tiles, trail, passed, 0, 'v'))
}

func getTileCount(point, end [2]int16, tiles [][]byte, trail []bool, passed [][]int16, steps int16, dir byte) int16 {
	if point[0] < 0 || point[0] >= int16(len(tiles[0])) || point[1] < 0 || point[1] >= int16(len(tiles)) {
		return 0
	}
	if tiles[point[1]][point[0]] != '.' && tiles[point[1]][point[0]] != dir {
		return 0
	}
	trailKey := (int(point[1]) * len(tiles[0])) + int(point[0])
	if trail[trailKey] {
		return 0
	}
	if steps > 0 && passed[point[1]][point[0]] >= steps {
		return 0
	}
	trail[trailKey] = true
	passed[point[1]][point[0]] = steps
	if point == end {
		return steps
	}
	if tiles[point[1]][point[0]] == dir {
		return getTileCount([2]int16{point[0] + dirs[dir][0], point[1] + dirs[dir][1]}, end, tiles, trail, passed, steps+1, dir)
	}

	s := steps
	for dir, mtx := range dirs {
		newPoint := [2]int16{point[0] + mtx[0], point[1] + mtx[1]}
		newTrail := slices.Clone(trail)
		newSteps := getTileCount(newPoint, end, tiles, newTrail, passed, steps+1, dir)
		if newSteps > s {
			s = newSteps
		}
	}

	return s
}
