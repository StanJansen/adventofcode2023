package step1

import (
	"math"
	"strconv"
	"strings"
)

type Direction rune

type Point struct {
	X, Y int
}

var directions = map[rune]Point{
	'U': {X: 0, Y: -1},
	'R': {X: 1, Y: 0},
	'D': {X: 0, Y: 1},
	'L': {X: -1, Y: 0},
}

func Solve(input string) int {
	lines := strings.Split(input, "\n")
	p := Point{X: 0, Y: 0}
	points := []Point{p}
	c := 2
	for _, line := range lines {
		dir := rune(line[0])
		v, _ := strconv.Atoi(line[2 : len(line)-10])
		c += v
		p.X += directions[dir].X * int(v)
		p.Y += directions[dir].Y * int(v)
		points = append(points, p)
	}

	return shoelace(points) + c/2
}

func shoelace(p []Point) int {
	n := len(p)
	area := p[n-1].X*p[0].Y - p[0].X*p[n-1].Y
	for i := 0; i < n-1; i++ {
		area += p[i].X*p[i+1].Y - p[i+1].X*p[i].Y
	}
	return int(math.Abs(float64(area)) / 2)
}
