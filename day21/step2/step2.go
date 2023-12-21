package step2

import (
	"bytes"
	"strings"
)

const STEPS = 26501365
const MAX_INT = 255

type Garden struct {
	Grid        [][]byte
	StartPoints [][4]byte
	Seen        map[[4]byte]bool
	Size        byte
}

func Solve(input string) int {
	garden := createGarden(input)
	size := int(garden.Size)
	subSteps := size - STEPS%(size+1)

	for i := 0; i < subSteps; i++ {
		garden.TakeStep()
	}
	a := len(garden.StartPoints)
	for i := 0; i <= size; i++ {
		garden.TakeStep()
	}
	b := len(garden.StartPoints)
	for i := 0; i <= size; i++ {
		garden.TakeStep()
	}
	c := len(garden.StartPoints)

	m := (STEPS - subSteps) / int(garden.Size+1)
	return (a/2-b+c/2)*m*m + (-3*(a/2)+2*b-c/2)*m + a // Lagrange multiplier
}

func createGarden(input string) Garden {
	grid := bytes.Fields([]byte(input))
	startIdx := strings.IndexRune(strings.ReplaceAll(input, "\n", ""), 'S')
	startX := startIdx % len(grid[0])

	return Garden{
		Grid:        grid,
		StartPoints: [][4]byte{{byte(startX), byte((startIdx - startX) / len(grid[0])), 0, 0}},
		Size:        byte(len(grid) - 1),
	}
}

func (g *Garden) TakeStep() {
	sp := [][4]byte{}
	g.Seen = make(map[[4]byte]bool, len(g.StartPoints)*4)

	for _, p := range g.StartPoints {
		newPoints := [4][4]byte{
			{p[0] - 1, p[1], p[2], p[3]},
			{p[0] + 1, p[1], p[2], p[3]},
			{p[0], p[1] - 1, p[2], p[3]},
			{p[0], p[1] + 1, p[2], p[3]},
		}
		for _, newPoint := range newPoints {
			if newPoint[0] == MAX_INT {
				newPoint[0] = g.Size
				newPoint[2]--
			} else if newPoint[0] > g.Size {
				newPoint[0] = 0
				newPoint[2]++
			} else if newPoint[1] == MAX_INT {
				newPoint[1] = g.Size
				newPoint[3]--
			} else if newPoint[1] > g.Size {
				newPoint[1] = 0
				newPoint[3]++
			}
			if g.Grid[newPoint[1]][newPoint[0]] == '#' {
				continue
			}
			if _, ok := g.Seen[newPoint]; ok {
				continue
			}
			sp = append(sp, newPoint)
			g.Seen[newPoint] = true
		}
	}

	g.StartPoints = sp
}
