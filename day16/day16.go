package day16

type Tiles [][]byte
type Direction int
type Visited map[int8]map[int8]map[Direction]bool

const (
	UP    = Direction(0)
	RIGHT = Direction(1)
	DOWN  = Direction(2)
	LEFT  = Direction(3)
)

func (t Tiles) Energized(startX, startY int8, d Direction) int {
	v := make(Visited, len(t))
	for y := range t {
		v[int8(y)] = make(map[int8]map[Direction]bool, len(t[y]))
		for x := range t[y] {
			v[int8(y)][int8(x)] = make(map[Direction]bool, 4)
		}
	}
	t.visit(startX, startY, d, v)

	sum := 0
	for y := range v {
		for x := range v[y] {
			for d := range v[y][x] {
				if v[y][x][d] {
					sum++
					break
				}
			}
		}
	}
	return sum
}

func (t Tiles) visit(x, y int8, d Direction, v Visited) {
	if x < 0 || y < 0 || y >= int8(len(t)) || x >= int8(len(t[y])) {
		return
	}
	if v[y][x][d] {
		return
	}
	v[y][x][d] = true

	switch t[y][x] {
	case '-':
		if d == UP || d == DOWN {
			t.visit(x-1, y, LEFT, v)
			t.visit(x+1, y, RIGHT, v)
			return
		}
	case '|':
		if d == LEFT || d == RIGHT {
			t.visit(x, y-1, UP, v)
			t.visit(x, y+1, DOWN, v)
			return
		}
	case '\\':
		switch d {
		case UP:
			d = LEFT
		case RIGHT:
			d = DOWN
		case DOWN:
			d = RIGHT
		case LEFT:
			d = UP
		}
	case '/':
		switch d {
		case UP:
			d = RIGHT
		case RIGHT:
			d = UP
		case DOWN:
			d = LEFT
		case LEFT:
			d = DOWN
		}
	}

	nextX, nextY := x, y
	switch d {
	case UP:
		nextY--
	case RIGHT:
		nextX++
	case DOWN:
		nextY++
	case LEFT:
		nextX--
	}

	t.visit(nextX, nextY, d, v)
}
