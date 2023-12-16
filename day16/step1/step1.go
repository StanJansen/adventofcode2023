package step1

import (
	"bytes"

	"github.com/StanJansen/adventofcode2023/day16"
)

func Solve(input string) int {
	tiles := day16.Tiles(bytes.Fields([]byte(input)))

	return tiles.Energized(0, 0, day16.RIGHT)
}
