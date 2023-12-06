package step1

import (
	"math"
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
		t, _ := strconv.ParseFloat(time, 10)
		d, _ := strconv.ParseFloat(distances[idx], 10)
		sq := math.Sqrt(t*t - 4*d)
		c := int(math.Abs(math.Ceil((-t-sq)/2) - math.Floor((-t+sq)/2)))
		if 0 == p {
			p = c
		} else {
			p *= c
		}
	}

	return p
}
