package step2

import (
	"bytes"
	"sync"

	"github.com/StanJansen/adventofcode2023/day16"
)

func Solve(input string) int {
	tiles := day16.Tiles(bytes.Fields([]byte(input)))
	sums := make([]int, len(tiles)*2+len(tiles[0])*2)
	wg := sync.WaitGroup{}
	wg.Add(len(sums))

	count := func(x, y, i int, d day16.Direction) {
		sums[i] = tiles.Energized(int8(x), int8(y), d)
		wg.Done()
	}

	i := 0
	for y := range tiles {
		go count(0, y, i, day16.RIGHT)
		i++
		go count(len(tiles[y])-1, y, i, day16.LEFT)
		i++
	}
	for x := range tiles[0] {
		go count(x, 0, i, day16.DOWN)
		i++
		go count(x, len(tiles)-1, i, day16.UP)
		i++
	}

	wg.Wait()

	highest := 0
	for _, sum := range sums {
		highest = max(highest, sum)
	}

	return highest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
