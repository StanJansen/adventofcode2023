package step2

import (
	"math"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

var directions = map[rune]Point{
	'3': {X: 0, Y: -1},
	'0': {X: 1, Y: 0},
	'1': {X: 0, Y: 1},
	'2': {X: -1, Y: 0},
}

func Solve(input string) int {
	lines := strings.Split(input, "\n")
	p := Point{X: 0, Y: 0}
	points := []Point{p}
	c := 2
	for _, line := range lines {
		dir := rune(line[len(line)-2])
		hex := line[len(line)-7 : len(line)-2]
		v, _ := strconv.ParseInt(hex, 16, 64)
		c += int(v)
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
