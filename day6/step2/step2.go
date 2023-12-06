package step2

import (
	"strings"

	"github.com/StanJansen/adventofcode2023/day6/step1"
)

func Solve(input string) int {
	return step1.Solve(strings.ReplaceAll(input, " ", ""))
}
