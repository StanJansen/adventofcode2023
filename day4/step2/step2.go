package step2

import (
	"strings"
)

func Solve(input string) int {
	lines := strings.Split(input, "\n")
	cards := make([]int, len(lines))
	sum := 0
	for i, line := range lines {
		parts := strings.Split(line, "|")
		nrs := " " + parts[0] + " "
		var wins int
		for _, n := range strings.Split(parts[1], " ") {
			if n != "" && strings.Contains(nrs, " "+n+" ") {
				wins++
				cards[i+wins] += cards[i] + 1
			}
		}
		sum += cards[i] + 1
	}
	return sum
}
