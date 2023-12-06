package step1

import (
	"strconv"
	"strings"
)

type Solver struct{}

var seeds []uint64
var mappings [][][3]uint64

func (Solver) Solve(input string) int {
	parseInput(input)

	source := seeds
	for _, mapping := range mappings {
		newSource := []uint64{}
		for _, s := range source {
			v := s
			for _, m := range mapping {
				if s >= m[1] && s < m[1]+m[2] {
					v = m[0] + s - m[1]
					break
				}
			}
			newSource = append(newSource, v)
		}
		source = newSource
	}

	min := source[0]

	for _, i := range source {
		if min > i {
			min = i
		}
	}

	return int(min)
}

func parseInput(input string) {
	lines := strings.Split(input+"\n", "\n")

	sLine := strings.Replace(lines[0], "seeds: ", "", 1)
	for _, seed := range strings.Split(sLine, " ") {
		s, _ := strconv.Atoi(seed)
		seeds = append(seeds, uint64(s))
	}

	var mapping [][3]uint64
	for _, line := range lines[2:] {
		if strings.Contains(line, ":") {
			mapping = [][3]uint64{}
			continue
		}

		if len(strings.TrimSpace(line)) == 0 {
			mappings = append(mappings, mapping)
			continue
		}

		digits := strings.Split(line, " ")
		var mappingLine [3]uint64
		for i, digit := range digits {
			d, _ := strconv.Atoi(digit)
			mappingLine[i] = uint64(d)
		}
		mapping = append(mapping, mappingLine)
	}
}
