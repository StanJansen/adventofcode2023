package step2

import (
	"bytes"
	"fmt"
	"strings"
)

const STEPS = 26501365

type Garden struct {
	Grid        [][]byte
	StartPoints [][5]int
	Seen        map[[5]int]bool
}

func Solve(input string) int {
	garden := createGarden(input)
	gridSize := len(garden.Grid)
	subSteps := gridSize - garden.StartPoints[0][0] - 1

	for i := 0; i < subSteps; i++ {
		garden.TakeStep()
	}
	a := len(garden.StartPoints)
	for i := 0; i < gridSize; i++ {
		garden.TakeStep()
	}
	b := len(garden.StartPoints)
	for i := 0; i < gridSize; i++ {
		garden.TakeStep()
	}
	c := len(garden.StartPoints)

	fmt.Println(a, b, c)

	m := (STEPS - subSteps) / gridSize
	return (a/2-b+c/2)*m*m + (-3*(a/2)+2*b-c/2)*m + a // Lagrange multiplier
}

func createGarden(input string) Garden {
	grid := bytes.Fields([]byte(input))
	startIdx := strings.IndexRune(strings.ReplaceAll(input, "\n", ""), 'S')
	startX := startIdx % len(grid[0])

	return Garden{
		Grid:        grid,
		StartPoints: [][5]int{{startX, (startIdx - startX) / len(grid[0]), STEPS, 0, 0}},
		Seen:        map[[5]int]bool{},
	}
}

func (g *Garden) TakeStep() {
	sp := [][5]int{}

	for _, p := range g.StartPoints {
		newPoints := [][5]int{
			{p[0] - 1, p[1], p[2] - 1, p[3], p[4]},
			{p[0] + 1, p[1], p[2] - 1, p[3], p[4]},
			{p[0], p[1] - 1, p[2] - 1, p[3], p[4]},
			{p[0], p[1] + 1, p[2] - 1, p[3], p[4]},
		}
		for _, newPoint := range newPoints {
			if newPoint[0] < 0 {
				newPoint[0] = len(g.Grid[0]) - 1
				newPoint[3]--
			} else if newPoint[0] > len(g.Grid[0])-1 {
				newPoint[0] = 0
				newPoint[3]++
			} else if newPoint[1] < 0 {
				newPoint[1] = len(g.Grid) - 1
				newPoint[4]--
			} else if newPoint[1] > len(g.Grid)-1 {
				newPoint[1] = 0
				newPoint[4]++
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
