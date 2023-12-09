package step1

import (
	"strconv"
	"strings"
)

func Solve(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0

	for _, line := range lines {
		sum += solveLine(line)
	}

	return sum
}

func solveLine(line string) int {
	var nrs []int
	for _, v := range strings.Split(line, " ") {
		nr, _ := strconv.Atoi(v)
		nrs = append(nrs, nr)
	}

	lastVals := []int{nrs[len(nrs)-1]}
	for {
		var row []int
		isLastRow := true
		for i := 0; i < len(nrs)-1; i++ {
			v := nrs[i+1] - nrs[i]
			row = append(row, v)
			if 0 != v {
				isLastRow = false
			}
		}

		lastVals = append(lastVals, row[len(row)-1])

		if isLastRow {
			break
		}

		nrs = row
	}

	result := 0
	for i := len(lastVals) - 1; i >= 0; i-- {
		result += lastVals[i]
	}

	return result
}
