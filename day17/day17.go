package day17

import (
	"bytes"
	"strconv"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

type Point struct {
	Y, X int
}

type Step struct {
	Point     Point
	Direction Direction
	Straight  byte
}

type Queued struct {
	Step     Step
	HeatLoss int
}

type Direction byte

const (
	UP    = Direction(1)
	DOWN  = Direction(2)
	LEFT  = Direction(3)
	RIGHT = Direction(4)
)

func Solve(input string, minStraight, maxStraight byte) int {
	grid := bytes.Fields([]byte(input))
	end := Point{len(grid) - 1, len(grid[0]) - 1}
	cities := make([][]int, len(grid))
	for y, row := range grid {
		cities[y] = make([]int, len(row))
		for x, v := range row {
			heatLoss, _ := strconv.Atoi(string(v))
			cities[y][x] = heatLoss
		}
	}

	pq := priorityqueue.NewWith(func(a, b any) int {
		return a.(Queued).HeatLoss - b.(Queued).HeatLoss
	})
	pq.Enqueue(Queued{Step{
		Point:     Point{0, 0},
		Direction: RIGHT,
		Straight:  1,
	}, 0})
	pq.Enqueue(Queued{Step{
		Point:     Point{0, 0},
		Direction: DOWN,
		Straight:  1,
	}, 0})
	cache := make(map[Step]int)

	for !pq.Empty() {
		q, _ := pq.Dequeue()
		queued := q.(Queued)
		step := queued.Step
		point := step.Point
		if point.Y < 0 || point.Y >= len(cities) || point.X < 0 || point.X >= len(cities[0]) {
			continue
		}

		hl := cities[point.Y][point.X] + queued.HeatLoss
		if point == end {
			return hl - cities[0][0]
		}

		if v, ok := cache[step]; ok {
			if v <= hl {
				continue
			}
		}
		cache[step] = hl

		if step.Straight >= minStraight {
			if step.Direction == UP || step.Direction == DOWN {
				pq.Enqueue(Queued{Step{
					Point:     Point{point.Y, point.X - 1},
					Direction: LEFT,
					Straight:  1,
				}, hl})
				pq.Enqueue(Queued{Step{
					Point:     Point{point.Y, point.X + 1},
					Direction: RIGHT,
					Straight:  1,
				}, hl})
			} else {
				pq.Enqueue(Queued{Step{
					Point:     Point{point.Y - 1, point.X},
					Direction: UP,
					Straight:  1,
				}, hl})
				pq.Enqueue(Queued{Step{
					Point:     Point{point.Y + 1, point.X},
					Direction: DOWN,
					Straight:  1,
				}, hl})
			}
		}

		if step.Straight < maxStraight {
			switch step.Direction {
			case UP:
				pq.Enqueue(Queued{Step{
					Point:     Point{point.Y - 1, point.X},
					Direction: UP,
					Straight:  step.Straight + 1,
				}, hl})
			case DOWN:
				pq.Enqueue(Queued{Step{
					Point:     Point{point.Y + 1, point.X},
					Direction: DOWN,
					Straight:  step.Straight + 1,
				}, hl})
			case LEFT:
				pq.Enqueue(Queued{Step{
					Point:     Point{point.Y, point.X - 1},
					Direction: LEFT,
					Straight:  step.Straight + 1,
				}, hl})
			case RIGHT:
				pq.Enqueue(Queued{Step{
					Point:     Point{point.Y, point.X + 1},
					Direction: RIGHT,
					Straight:  step.Straight + 1,
				}, hl})
			}
		}
	}

	return -1
}
