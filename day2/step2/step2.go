package step2

import (
	"regexp"
	"strconv"
	"strings"
)

var colors = []string{"red", "green", "blue"}

func Solve(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		pow := 0
		for _, color := range colors {
			r := regexp.MustCompile("(\\d+)( " + color + ")")
			m := r.FindAllString(line, -1)

			l := 0
			for _, v := range m {
				v = strings.Replace(v, " "+color, "", -1)
				if amount, _ := strconv.Atoi(v); amount > l {
					l = amount
				}
			}

			if pow == 0 {
				pow = l
			} else {
				pow *= l
			}
		}

		sum += pow
	}

	return sum
}
