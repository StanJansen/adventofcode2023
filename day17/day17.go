package day17

import (
	"bytes"
	"strconv"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

type Map struct {
	Cities [][]int
	PQ     *priorityqueue.Queue
}

type Point struct {
	Y, X int
}

type Step struct {
	Point     Point
	Direction Direction
	Straight  byte
}

type Queued struct {
	Step          Step
	TotalHeatLoss int
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
	m := CreateMap(grid)

	return m.FindPath(Point{len(grid) - 1, len(grid[0]) - 1}, minStraight, maxStraight)
}

func CreateMap(grid [][]byte) Map {
	cities := make([][]int, len(grid))
	for y, row := range grid {
		cities[y] = make([]int, len(row))
		for x, v := range row {
			heatLoss, _ := strconv.Atoi(string(v))
			cities[y][x] = heatLoss
		}
	}

	return Map{
		Cities: cities,
		PQ:     InitPQ(),
	}
}

func InitPQ() *priorityqueue.Queue {
	pq := priorityqueue.NewWith(func(a, b any) int {
		return a.(Queued).TotalHeatLoss - b.(Queued).TotalHeatLoss
	})

	// Start from top left, go both right and down
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

	return pq
}

func (m Map) FindPath(end Point, minStraight, maxStraight byte) int {
	cache := make(map[Step]int)
	for {
		q, _ := m.PQ.Dequeue()
		queued := q.(Queued)
		step := queued.Step
		point := step.Point
		if point.Y < 0 || point.Y >= len(m.Cities) || point.X < 0 || point.X >= len(m.Cities[0]) {
			// Out of bounds
			continue
		}

		// Add the heat loss of the current point to the total
		hl := m.Cities[point.Y][point.X] + queued.TotalHeatLoss
		if point == end {
			if step.Straight < minStraight {
				// Must end with the min straight
				continue
			}
			return hl - m.Cities[0][0]
		}

		// Check cache
		if v, ok := cache[step]; ok {
			if v <= hl {
				// Less efficient than a cached path
				continue
			}
		}
		cache[step] = hl

		// Only allow a straight path if it's below the max
		if step.Straight < maxStraight {
			step.EnqueueStraight(m.PQ, hl)
		}

		// Only allow a turn if it has been going straight for at least the min
		if step.Straight >= minStraight {
			step.EnqueueTurns(m.PQ, hl)
		}
	}
}

func (s Step) EnqueueStraight(pq *priorityqueue.Queue, hl int) {
	p := s.Point
	switch s.Direction {
	case UP:
		pq.Enqueue(Queued{Step{
			Point:     Point{p.Y - 1, p.X},
			Direction: UP,
			Straight:  s.Straight + 1,
		}, hl})
	case DOWN:
		pq.Enqueue(Queued{Step{
			Point:     Point{p.Y + 1, p.X},
			Direction: DOWN,
			Straight:  s.Straight + 1,
		}, hl})
	case LEFT:
		pq.Enqueue(Queued{Step{
			Point:     Point{p.Y, p.X - 1},
			Direction: LEFT,
			Straight:  s.Straight + 1,
		}, hl})
	case RIGHT:
		pq.Enqueue(Queued{Step{
			Point:     Point{p.Y, p.X + 1},
			Direction: RIGHT,
			Straight:  s.Straight + 1,
		}, hl})
	}
}

func (s Step) EnqueueTurns(pq *priorityqueue.Queue, hl int) {
	p := s.Point
	if s.Direction == UP || s.Direction == DOWN {
		pq.Enqueue(Queued{Step{
			Point:     Point{p.Y, p.X - 1},
			Direction: LEFT,
			Straight:  1,
		}, hl})
		pq.Enqueue(Queued{Step{
			Point:     Point{p.Y, p.X + 1},
			Direction: RIGHT,
			Straight:  1,
		}, hl})
	} else {
		pq.Enqueue(Queued{Step{
			Point:     Point{p.Y - 1, p.X},
			Direction: UP,
			Straight:  1,
		}, hl})
		pq.Enqueue(Queued{Step{
			Point:     Point{p.Y + 1, p.X},
			Direction: DOWN,
			Straight:  1,
		}, hl})
	}
}
