package step2

import (
	"bytes"
)

type Platform [][]byte
type Direction byte

const (
	SOLID = byte('#')
	ROUND = byte('O')
	EMPTY = byte('.')

	NORTH = Direction(1)
	WEST  = Direction(2)
	SOUTH = Direction(3)
	EAST  = Direction(4)
)

func Solve(input string) int {
	p := Platform(bytes.Fields([]byte(input)))
	hashes := map[string]int{}
	len := 1000000000
	for i := 0; i < len; i++ {
		p.Tilt(NORTH)
		p.Tilt(WEST)
		p.Tilt(SOUTH)
		p.Tilt(EAST)
		h := bytes.Join(p, []byte{})
		if _, ok := hashes[string(h)]; ok {
			i = len - (len-i)%(i-hashes[string(h)])
		}
		hashes[string(h)] = i
	}
	return p.Load()
}

func (p Platform) Tilt(d Direction) {
	switch d {
	case NORTH:
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
	case WEST:
		for y := 0; y < len(p); y++ {
			available := 0
			for x := 0; x < len(p[y]); x++ {
				switch p[y][x] {
				case SOLID:
					available = x + 1
				case ROUND:
					if available < x {
						p[y][available] = ROUND
						p[y][x] = EMPTY
					}
					available++
				}
			}
		}
	case SOUTH:
		for x := 0; x < len(p[0]); x++ {
			available := len(p) - 1
			for y := len(p) - 1; y >= 0; y-- {
				switch p[y][x] {
				case SOLID:
					available = y - 1
				case ROUND:
					if available > y {
						p[available][x] = ROUND
						p[y][x] = EMPTY
					}
					available--
				}
			}
		}
	case EAST:
		for y := 0; y < len(p); y++ {
			available := len(p[y]) - 1
			for x := len(p[y]) - 1; x >= 0; x-- {
				switch p[y][x] {
				case SOLID:
					available = x - 1
				case ROUND:
					if available > x {
						p[y][available] = ROUND
						p[y][x] = EMPTY
					}
					available--
				}
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
