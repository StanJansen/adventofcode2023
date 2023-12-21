package step1

import (
	"bytes"
	"strings"
)

const STEPS = 64

type Garden struct {
	Grid        [][]byte
	StartPoints [][3]int
	Seen        map[[3]int]bool
}

func Solve(input string) int {
	garden := createGarden(input)
	for i := 0; i < STEPS; i++ {
		garden.TakeStep()
	}

	return len(garden.StartPoints)
}

func createGarden(input string) Garden {
	grid := bytes.Fields([]byte(input))
	startIdx := strings.IndexRune(strings.ReplaceAll(input, "\n", ""), 'S')
	startX := startIdx % len(grid[0])

	return Garden{
		Grid:        grid,
		StartPoints: [][3]int{{startX, (startIdx - startX) / len(grid[0]), STEPS}},
		Seen:        map[[3]int]bool{},
	}
}

func (g *Garden) TakeStep() {
	sp := [][3]int{}

	for _, p := range g.StartPoints {
		newPoints := [][3]int{
			{p[0] - 1, p[1], p[2] - 1},
			{p[0] + 1, p[1], p[2] - 1},
			{p[0], p[1] - 1, p[2] - 1},
			{p[0], p[1] + 1, p[2] - 1},
		}
		for _, newPoint := range newPoints {
			if newPoint[0] < 0 || newPoint[0] > len(g.Grid[0])-1 || newPoint[1] < 0 || newPoint[1] > len(g.Grid)-1 {
				continue
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
