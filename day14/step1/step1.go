package step1

import "bytes"

type Platform [][]byte

const (
	SOLID = byte('#')
	ROUND = byte('O')
	EMPTY = byte('.')
)

func Solve(input string) int {
	p := Platform(bytes.Fields([]byte(input)))
	p.Tilt()
	return p.Load()
}

func (p Platform) Tilt() {
	for x := 0; x < len(p[0]); x++ {
		available := 0
		for y := 0; y < len(p); y++ {
			switch p[y][x] {
			case SOLID:
				available = y + 1
			case ROUND:
				if available < y {
					p[available][x] = ROUND
					p[y][x] = EMPTY
				}
				available++
			}
		}
	}
}

func (p Platform) Load() int {
	sum := 0
	for x := 0; x < len(p); x++ {
		for y := 0; y < len(p[x]); y++ {
			if p[y][x] == ROUND {
				sum += len(p) - int(y)
			}
		}
	}

	return sum
}
