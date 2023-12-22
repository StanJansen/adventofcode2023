package step1

import (
	"fmt"
	"sort"
	"strings"
)

type Point struct {
	X, Y, Z int
}
type Brick struct {
	Start, End Point
}
type Bricks []Brick

func Solve(input string) int {
	bricks := getBricks(input)

	below, above := bricks.GetIndexes()

	sum := len(bricks)
	for i := 0; i < len(bricks); i++ {
		for _, idx := range above[i] {
			if below[idx] <= 1 {
				sum--
				break
			}
		}
	}

	return sum
}

func getBricks(input string) Bricks {
	lines := strings.Split(input, "\n")
	bricks := make(Bricks, 0, len(lines))
	for _, line := range lines {
		b := Brick{}
		fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &b.Start.X, &b.Start.Y, &b.Start.Z, &b.End.X, &b.End.Y, &b.End.Z)
		bricks = append(bricks, b)
	}
	sort.Slice(bricks, func(a, b int) bool {
		return bricks[a].Start.Z < bricks[b].Start.Z
	})
	return bricks
}

func (bricks Bricks) GetIndexes() ([]int, [][]int) {
	below := make([]int, len(bricks))
	above := make([][]int, len(bricks))
	idx := 0
	for idx < len(bricks) {
		for i := idx; i < len(bricks); i++ {
			if bricks[i].Start.Z == 1 {
				idx++
				continue
			}

			collisions := bricks[i].Descend(bricks[:i])
			if len(collisions) == 0 {
				break
			} else {
				below[i] = len(collisions)
				for _, brick := range collisions {
					above[brick] = append(above[brick], i)
				}
				idx++
			}
		}
	}
	return below, above
}

func (brick *Brick) Descend(nextBricks []Brick) []int {
	brick.Start.Z--
	brick.End.Z--
	collisions := []int{}
	for i, b := range nextBricks {
		if brick.Overlaps(b) {
			collisions = append(collisions, i)
		}
	}
	if len(collisions) > 0 {
		brick.Start.Z++
		brick.End.Z++
	}
	return collisions
}

func (a *Brick) Overlaps(b Brick) bool {
	return (a.Start.X <= b.End.X && a.End.X >= b.Start.X) &&
		(a.Start.Y <= b.End.Y && a.End.Y >= b.Start.Y) &&
		(a.Start.Z <= b.End.Z && a.End.Z >= b.Start.Z)
}
