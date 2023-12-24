package step1

import (
	"fmt"
	"strings"
)

type Point struct {
	X, Y, Z float64
}

type Hail struct {
	Position, Velocity Point
}

const MIN float64 = 200000000000000
const MAX float64 = 400000000000000

func Solve(input string) int {
	hails := parseInput(input)

	sum := 0
	for i, a := range hails {
		for _, b := range hails[i+1:] {
			if a.Collides(b) {
				sum++
			}
		}
	}

	return sum
}

func parseInput(input string) []Hail {
	lines := strings.Split(input, "\n")
	hails := make([]Hail, len(lines))

	for i, line := range lines {
		h := Hail{}
		fmt.Sscanf(line, "%f, %f, %f @ %f, %f, %f", &h.Position.X, &h.Position.Y, &h.Position.Z, &h.Velocity.X, &h.Velocity.Y, &h.Velocity.Z)
		hails[i] = h
	}

	return hails
}

func (a Hail) Collides(b Hail) bool {
	v1 := a.Velocity.Y / a.Velocity.X
	v2 := b.Velocity.Y / b.Velocity.X
	if v1 == v2 {
		return false
	}

	x := ((v2 * b.Position.X) - (v1 * a.Position.X) + a.Position.Y - b.Position.Y) / (v2 - v1)
	if x < MIN || x > MAX {
		return false
	}

	y := (v1 * (x - a.Position.X)) + a.Position.Y
	if y < MIN || y > MAX {
		return false
	}

	return a.Supports(x, y) && b.Supports(x, y)
}

func (hail Hail) Supports(x, y float64) bool {
	return !((x > hail.Position.X && 0 > hail.Velocity.X) ||
		(x < hail.Position.X && 0 < hail.Velocity.X) ||
		(y > hail.Position.Y && 0 > hail.Velocity.Y) ||
		(y < hail.Position.Y && 0 < hail.Velocity.Y))
}
