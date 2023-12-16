package day16

type Tiles [][]byte
type Direction int
type Visited [][][4]bool

const (
	UP    = Direction(0)
	RIGHT = Direction(1)
	DOWN  = Direction(2)
	LEFT  = Direction(3)
)

func (t Tiles) Energized(startX, startY int8, d Direction) int {
	v := make(Visited, len(t))
	for y := range v {
		v[y] = make([][4]bool, len(t[y]))
	}
	return t.visit(startX, startY, d, v)
}

func (t Tiles) visit(x, y int8, d Direction, v Visited) (visited int) {
	if x < 0 || y < 0 || y >= int8(len(t)) || x >= int8(len(t[y])) {
		return
	}
	if v[y][x][d] {
		return
	}

	visited++
	for _, b := range v[y][x] {
		if b {
			visited--
			break
		}
	}
	v[y][x][d] = true

	switch t[y][x] {
	case '-':
		if d == UP || d == DOWN {
			visited += t.visit(x-1, y, LEFT, v)
			visited += t.visit(x+1, y, RIGHT, v)
			return
		}
	case '|':
		if d == LEFT || d == RIGHT {
			visited += t.visit(x, y-1, UP, v)
			visited += t.visit(x, y+1, DOWN, v)
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

	switch d {
	case UP:
		y--
	case RIGHT:
		x++
	case DOWN:
		y++
	case LEFT:
		x--
	}

	return visited + t.visit(x, y, d, v)
}
