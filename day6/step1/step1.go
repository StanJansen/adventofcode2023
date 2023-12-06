package step1

import (
	"regexp"
	"strconv"
	"strings"
)

var regex = regexp.MustCompile("\\d+")

func Solve(input string) int {
	lines := strings.Split(input, "\n")
	times := regex.FindAllString(lines[0], -1)
	distances := regex.FindAllString(lines[1], -1)

	p := 0
	for idx, time := range times {
		t, _ := strconv.Atoi(time)
		d, _ := strconv.Atoi(distances[idx])

		c := 0
		for i := 1; i < t; i++ {
			if 0 > i*i-t*i+d {
				c++
			}
		}

		if 0 == p {
			p = c
		} else {
			p *= c
		}
	}

	return p
}
