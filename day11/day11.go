package day11

import (
	"strings"
)

type Point struct {
	X int
	Y int
}

func SolveForMultiplier(input string, multiplier int) int {
	lines := strings.Split(input, "\n")

	xWithoutPoint := make([]bool, len(lines[0]))
x:
	for x := 0; x < len(lines[0]); x++ {
		for y := 0; y < len(lines); y++ {
			if lines[y][x] == '#' {
				continue x
			}
		}
		xWithoutPoint[x] = true
	}

	points := make([]Point, strings.Count(input, "#"))
	var point, extraY int
	for y, line := range lines {
		if !strings.Contains(line, "#") {
			extraY += multiplier - 1
			continue
		}

		var extraX int
		for x, char := range line {
			if char == '#' {
				points[point] = Point{X: x + extraX, Y: y + extraY}
				point++
			} else if xWithoutPoint[x] {
				extraX += multiplier - 1
			}
		}
	}

	var sum int
	for i, point := range points {
		for j := i + 1; j < len(points); j++ {
			sum += diff(point.X, points[j].X) + points[j].Y - point.Y
		}
	}

	return sum
}

func diff(a int, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
