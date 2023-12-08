package step1

import (
	"strings"
)

func Solve(input string) int {
	lines := strings.Split(input, "\n")

	m := make(map[string][2]string, len(lines)-2)
	for _, line := range lines[2:] {
		m[line[0:3]] = [2]string{line[7:10], line[12:15]}
	}

	k := "AAA"
	var s int
	for k != "ZZZ" {
		for _, i := range lines[0] {
			s++
			if i == 'L' {
				k = m[k][0]
			} else {
				k = m[k][1]
			}
		}
	}

	return s
}
