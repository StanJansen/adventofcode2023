package step1

import (
	"regexp"
	"strconv"
	"strings"
)

var regex = regexp.MustCompile("[0-9]")

func Solve(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		matches := regex.FindAllString(line, -1)
		amount, _ := strconv.Atoi(matches[0] + matches[len(matches)-1])
		sum += amount
	}

	return sum
}
