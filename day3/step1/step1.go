package step1

import (
	"regexp"
	"strconv"
	"strings"
)

type Solver struct{}

var dRegex = regexp.MustCompile("\\d+")
var sRegex = regexp.MustCompile("[^0-9.]")

func (Solver) Solve(input string) int {
	lines := strings.Split(input, "\n")

	sIdxs := map[int][][]int{}
	for idx, line := range lines {
		sIdxs[idx] = sRegex.FindAllStringIndex(line, -1)
	}

	sum := 0
	for idx, line := range lines {
		dIdxs := dRegex.FindAllStringIndex(line, -1)
	dig:
		for _, dIdx := range dIdxs {
			amount, _ := strconv.Atoi(line[dIdx[0]:dIdx[1]])
			if idx > 0 {
				for _, sIdx := range sIdxs[idx-1] {
					if sIdx[0] >= dIdx[0]-1 && sIdx[0] <= dIdx[1] {
						sum += amount
						continue dig
					}
				}
			}
			for _, sIdx := range sIdxs[idx] {
				if sIdx[0] == dIdx[0]-1 || sIdx[0] == dIdx[1] {
					sum += amount
					continue dig
				}
			}
			if idx < len(lines)-1 {
				for _, sIdx := range sIdxs[idx+1] {
					if sIdx[0] >= dIdx[0]-1 && sIdx[0] <= dIdx[1] {
						sum += amount
						continue dig
					}
				}
			}
		}
	}

	return sum
}
