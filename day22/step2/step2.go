package step2

import (
	"fmt"
	"sort"
	"strings"

	"github.com/emirpasic/gods/queues/arrayqueue"
)

type Point struct {
	X, Y, Z int
}
type Brick struct {
	Start, End Point
}
type Bricks []Brick
type BrickIndexes [][]int

func Solve(input string) int {
	bricks := getBricks(input)

	below, above := bricks.GetIndexes()

	isStuck := make([]bool, len(bricks))
	for i := 0; i < len(bricks); i++ {
		for _, idx := range above[i] {
			if len(below[idx]) <= 1 {
				isStuck[i] = true
				break
			}
		}
	}

	sum := 0
	queue := arrayqueue.New()
	for i := 0; i < len(bricks); i++ {
		if !isStuck[i] {
			continue
		}

		queue.Enqueue(i)
		removed := make([]bool, len(bricks))
		removed[i] = true
		for !queue.Empty() {
			q, _ := queue.Dequeue()
			for _, above := range above[q.(int)] {
				count := 0
				for _, below := range below[above] {
					if removed[below] {
						count++
					}
				}
				if len(below[above]) == count && !removed[above] {
					queue.Enqueue(above)
					removed[above] = true
					sum++
				}
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

func (bricks Bricks) GetIndexes() (BrickIndexes, BrickIndexes) {
	below := make(BrickIndexes, len(bricks))
	above := make(BrickIndexes, len(bricks))
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
				below[i] = collisions
				for _, brick := range collisions {
					above[brick] = append(above[brick], i)
				}
				idx++
			}
		}
	}
	return below, above
}

func (brick *Brick) Descend(nextBricks []Brick) (collisions []int) {
	brick.Start.Z--
	brick.End.Z--
	for i, b := range nextBricks {
		if brick.Overlaps(b) {
			collisions = append(collisions, i)
		}
	}
	if len(collisions) > 0 {
		brick.Start.Z++
		brick.End.Z++
	}
	return
}

func (a *Brick) Overlaps(b Brick) bool {
	return (a.Start.X <= b.End.X && a.End.X >= b.Start.X) &&
		(a.Start.Y <= b.End.Y && a.End.Y >= b.Start.Y) &&
		(a.Start.Z <= b.End.Z && a.End.Z >= b.Start.Z)
}
