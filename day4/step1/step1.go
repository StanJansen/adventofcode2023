package step1

import (
	"math"
	"strings"
)

func Solve(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, "|")
		nrs := " " + parts[0] + " "
		power := 0
		for _, n := range strings.Split(parts[1], " ") {
			if n != "" && strings.Contains(nrs, " "+n+" ") {
				power++
			}
		}
		if power <= 1 {
			sum += power
		} else {
			sum += int(math.Pow(2, float64(power-1)))
		}
	}
	return sum
}
