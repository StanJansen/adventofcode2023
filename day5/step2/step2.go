package step2

import (
	"strconv"
	"strings"
)

var seedParts []string
var mappings [][][3]uint64

func Solve(input string) uint64 {
	parseInput(input)

	lowest := uint64(0)
	for i := 0; i < len(seedParts); i += 2 {
		s1, _ := strconv.Atoi(seedParts[i])
		s2, _ := strconv.Atoi(seedParts[i+1])
		for j := s1; j < s1+s2; j++ {
			r := solveSeed(uint64(j))
			if lowest == 0 || lowest > r {
				lowest = r
			}
		}
	}

	return lowest
}

func solveSeed(seed uint64) uint64 {
	for _, mapping := range mappings {
		for _, m := range mapping {
			if seed >= m[1] && seed < m[1]+m[2] {
				seed = m[0] + seed - m[1]
				break
			}
		}
	}

	return seed
}

func parseInput(input string) {
	lines := strings.Split(input+"\n", "\n")

	sLine := strings.Replace(lines[0], "seeds: ", "", 1)
	seedParts = strings.Split(sLine, " ")

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
