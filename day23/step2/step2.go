package step2

import (
	"bytes"
	"slices"
)

var dirs = [4][2]int16{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func Solve(input string) int {
	tiles := bytes.Fields([]byte(input))
	passed := make([]int16, len(tiles)*len(tiles[0]))
	trail := make([]bool, len(tiles)*len(tiles[0]))
	point := [2]int16{1, 0}
	end := [2]int16{int16(len(tiles[0]) - 2), int16(len(tiles) - 1)}

	v, _ := getTileCount(point, end, tiles, trail, passed, 0)

	return int(v)
}

func getTileCount(point, end [2]int16, tiles [][]byte, trail []bool, passed []int16, steps int16) (int16, bool) {
	if point[0] < 0 || point[0] >= int16(len(tiles[0])) || point[1] < 0 || point[1] >= int16(len(tiles)) {
		return 0, false
	}
	if tiles[point[1]][point[0]] == '#' {
		return 0, false
	}
	trailKey := (int(point[1]) * len(tiles[0])) + int(point[0])
	if trail[trailKey] {
		return 0, false
	}
	if steps > 0 && passed[trailKey] >= steps {
		return 0, false
	}
	trail[trailKey] = true
	if point == end {
		passed[trailKey] = steps
		return steps, true
	}

	s := steps
	tOk := false
	for _, dir := range dirs {
		newPoint := [2]int16{point[0] + dir[0], point[1] + dir[1]}
		newTrail := slices.Clone(trail)
		newSteps, ok := getTileCount(newPoint, end, tiles, newTrail, passed, steps+1)
		if ok && newSteps > s {
			s = newSteps
			tOk = ok
		}
	}

	return s, tOk
}
