package step1

import (
	"bytes"
	"strings"
)

const STEPS = 64

type Garden struct {
	Grid        [][]byte
	StartPoints [][2]byte
	Seen        map[[2]byte]bool
	Size        byte
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
		StartPoints: [][2]byte{{byte(startX), byte((startIdx - startX) / len(grid[0]))}},
		Size:        byte(len(grid) - 1),
	}
}

func (g *Garden) TakeStep() {
	sp := [][2]byte{}
	g.Seen = make(map[[2]byte]bool, len(g.StartPoints)*4)

	for _, p := range g.StartPoints {
		newPoints := [][2]byte{
			{p[0] - 1, p[1]},
			{p[0] + 1, p[1]},
			{p[0], p[1] - 1},
			{p[0], p[1] + 1},
		}
		for _, newPoint := range newPoints {
			if newPoint[0] < 0 || newPoint[0] > g.Size || newPoint[1] < 0 || newPoint[1] > g.Size {
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
