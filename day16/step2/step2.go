package step2

import (
	"bytes"

	"github.com/StanJansen/adventofcode2023/day16"
)

func Solve(input string) int {
	tiles := day16.Tiles(bytes.Fields([]byte(input)))
	highest := 0
	for y := range tiles {
		highest = max(highest, tiles.Energized(0, int8(y), day16.RIGHT))
		highest = max(highest, tiles.Energized(int8(len(tiles[0])-1), int8(y), day16.LEFT))
	}
	for x := range tiles[0] {
		highest = max(highest, tiles.Energized(int8(x), 0, day16.DOWN))
		highest = max(highest, tiles.Energized(int8(x), int8(len(tiles)-1), day16.UP))
	}
	return highest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
