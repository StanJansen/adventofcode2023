package step2

import (
	"regexp"
	"strconv"
	"strings"
)

var dRegex = regexp.MustCompile("\\d+")
var sRegex = regexp.MustCompile("\\*")

func Solve(input string) int {
	lines := strings.Split(input, "\n")

	sIdxs := map[int][][]int{}
	for idx, line := range lines {
		sIdxs[idx] = sRegex.FindAllStringIndex(line, -1)
	}

	sResults := map[string][]int{}
	appendResult := func(lIdx int, sIdx int, amount int) {
		idx := strconv.Itoa(lIdx) + "-" + strconv.Itoa(sIdx)
		if _, ok := sResults[idx]; ok {
			sResults[idx] = append(sResults[idx], amount)
		} else {
			sResults[idx] = []int{amount}
		}
	}

	for idx, line := range lines {
		dIdxs := dRegex.FindAllStringIndex(line, -1)
		for _, dIdx := range dIdxs {
			amount, _ := strconv.Atoi(line[dIdx[0]:dIdx[1]])
			if idx > 0 {
				for _, sIdx := range sIdxs[idx-1] {
					if sIdx[0] >= dIdx[0]-1 && sIdx[0] <= dIdx[1] {
						appendResult(idx-1, sIdx[0], amount)
					}
				}
			}
			for _, sIdx := range sIdxs[idx] {
				if sIdx[0] == dIdx[0]-1 || sIdx[0] == dIdx[1] {
					appendResult(idx, sIdx[0], amount)
				}
			}
			if idx < len(lines)-1 {
				for _, sIdx := range sIdxs[idx+1] {
					if sIdx[0] >= dIdx[0]-1 && sIdx[0] <= dIdx[1] {
						appendResult(idx+1, sIdx[0], amount)
					}
				}
			}
		}
	}

	sum := 0
	for _, sResult := range sResults {
		if len(sResult) == 2 {
			sum += sResult[0] * sResult[1]
		}
	}

	return sum
}
