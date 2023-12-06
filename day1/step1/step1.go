package step1

import (
	"regexp"
	"strconv"
	"strings"
)

type Solver struct{}

var regex = regexp.MustCompile("[0-9]")

func (Solver) Solve(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		matches := regex.FindAllString(line, -1)
		amount, _ := strconv.Atoi(matches[0] + matches[len(matches)-1])
		sum += amount
	}

	return sum
}
