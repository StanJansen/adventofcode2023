package step2

import (
	"fmt"
	"math"
	"strings"

	"gonum.org/v1/gonum/mat"
)

type Point struct {
	X, Y, Z float64
}

type Hail struct {
	Position, Velocity Point
}

func Solve(input string) int {
	hails := parseInput(input)
	h1 := hails[0]
	h2 := hails[1]
	h3 := hails[2]

	a := mat.NewDense(6, 6, []float64{
		-(h1.Velocity.Y - h2.Velocity.Y), h1.Velocity.X - h2.Velocity.X, 0, h1.Position.Y - h2.Position.Y, -(h1.Position.X - h2.Position.X), 0,
		-(h1.Velocity.Y - h3.Velocity.Y), h1.Velocity.X - h3.Velocity.X, 0, h1.Position.Y - h3.Position.Y, -(h1.Position.X - h3.Position.X), 0,

		0, -(h1.Velocity.Z - h2.Velocity.Z), h1.Velocity.Y - h2.Velocity.Y, 0, h1.Position.Z - h2.Position.Z, -(h1.Position.Y - h2.Position.Y),
		0, -(h1.Velocity.Z - h3.Velocity.Z), h1.Velocity.Y - h3.Velocity.Y, 0, h1.Position.Z - h3.Position.Z, -(h1.Position.Y - h3.Position.Y),

		-(h1.Velocity.Z - h2.Velocity.Z), 0, h1.Velocity.X - h2.Velocity.X, h1.Position.Z - h2.Position.Z, 0, -(h1.Position.X - h2.Position.X),
		-(h1.Velocity.Z - h3.Velocity.Z), 0, h1.Velocity.X - h3.Velocity.X, h1.Position.Z - h3.Position.Z, 0, -(h1.Position.X - h3.Position.X),
	})

	b := mat.NewDense(6, 1, []float64{
		(h1.Position.Y*h1.Velocity.X - h2.Position.Y*h2.Velocity.X) - (h1.Position.X*h1.Velocity.Y - h2.Position.X*h2.Velocity.Y),
		(h1.Position.Y*h1.Velocity.X - h3.Position.Y*h3.Velocity.X) - (h1.Position.X*h1.Velocity.Y - h3.Position.X*h3.Velocity.Y),

		(h1.Position.Z*h1.Velocity.Y - h2.Position.Z*h2.Velocity.Y) - (h1.Position.Y*h1.Velocity.Z - h2.Position.Y*h2.Velocity.Z),
		(h1.Position.Z*h1.Velocity.Y - h3.Position.Z*h3.Velocity.Y) - (h1.Position.Y*h1.Velocity.Z - h3.Position.Y*h3.Velocity.Z),

		(h1.Position.Z*h1.Velocity.X - h2.Position.Z*h2.Velocity.X) - (h1.Position.X*h1.Velocity.Z - h2.Position.X*h2.Velocity.Z),
		(h1.Position.Z*h1.Velocity.X - h3.Position.Z*h3.Velocity.X) - (h1.Position.X*h1.Velocity.Z - h3.Position.X*h3.Velocity.Z),
	})

	var x mat.Dense
	x.Solve(a, b)
	return int(math.Ceil(x.At(0, 0) + x.At(1, 0) + x.At(2, 0)))
}

func parseInput(input string) [3]Hail {
	lines := strings.Split(input, "\n")
	hails := [3]Hail{}

	for i, line := range lines[:3] {
		h := Hail{}
		fmt.Sscanf(line, "%f, %f, %f @ %f, %f, %f", &h.Position.X, &h.Position.Y, &h.Position.Z, &h.Velocity.X, &h.Velocity.Y, &h.Velocity.Z)
		hails[i] = h
	}

	return hails
}
