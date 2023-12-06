package step2

import (
	"strconv"
	"strings"
)

var seeds [][2]uint64
var mappings [][][3]uint64

func Solve(input string) int {
	parseInput(input)

	source := seeds
	for _, mapping := range mappings {
		newSource := [][2]uint64{}
		for _, s := range source {
			vs := applyMapping(s, mapping)

			newSource = append(newSource, vs...)
		}
		source = newSource
	}

	min := uint64(source[0][0])

	for _, i := range source {
		if 0 != i[0] && min > i[0] {
			min = i[0]
		}
	}

	return int(min)
}

func applyMapping(s [2]uint64, mapping [][3]uint64) [][2]uint64 {
	vs := [][2]uint64{}

	for _, m := range mapping {
		if (s[0] >= m[1] && s[0] < m[1]+m[2]) || (s[1] >= m[1] && s[1] < m[1]+m[2]) || (s[0] < m[1] && s[1] >= m[1]+m[2]) {
			mi := max([]uint64{s[0], m[1]})
			ma := min([]uint64{s[1], m[1] + m[2] - 1})
			if mi > s[0] {
				vs = append(vs, applyMapping([2]uint64{s[0], mi - 1}, mapping)...)
			}
			vs = append(vs, [2]uint64{
				m[0] + (mi - m[1]),
				m[0] + (ma - m[1]),
			})
			if ma < s[1] {
				vs = append(vs, applyMapping([2]uint64{ma + 1, s[1]}, mapping)...)
			}
			break
		}
	}

	if len(vs) == 0 {
		vs = append(vs, s)
	}

	return vs
}

func parseInput(input string) {
	lines := strings.Split(input+"\n", "\n")

	sLine := strings.Replace(lines[0], "seeds: ", "", 1)
	sParts := strings.Split(sLine, " ")
	for i := 0; i < len(sParts); i += 2 {
		s1, _ := strconv.Atoi(sParts[i])
		s2, _ := strconv.Atoi(sParts[i+1])
		seeds = append(seeds, [2]uint64{uint64(s1), uint64(s1 + s2 - 1)})
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

func min(source []uint64) uint64 {
	min := source[0]

	for _, i := range source {
		if min > i {
			min = i
		}
	}

	return min
}

func max(source []uint64) uint64 {
	max := source[0]

	for _, i := range source {
		if max < i {
			max = i
		}
	}

	return max
}
