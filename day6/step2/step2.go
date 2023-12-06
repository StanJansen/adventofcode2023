package step2

import (
	"strings"

	"github.com/StanJansen/adventofcode2023/day6/step1"
)

type Solver struct{}

func (Solver) Solve(input string) int {
	return step1.Solver{}.Solve(strings.ReplaceAll(input, " ", ""))
}
