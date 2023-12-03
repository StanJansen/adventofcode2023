package step1

import (
	"regexp"
	"strconv"
	"strings"
)

var regexes = [3]*regexp.Regexp{
	regexp.MustCompile("(Game \\d{1,}).*([2-9][0-9]|1[3-9]|\\d{3,})( red).*\n"),
	regexp.MustCompile("(Game \\d{1,}).*([2-9][0-9]|1[4-9]|\\d{3,})( green).*\n"),
	regexp.MustCompile("(Game \\d{1,}).*([2-9][0-9]|1[5-9]|\\d{3,})( blue).*\n"),
}
var gamer = regexp.MustCompile("Game (\\d{1,})")

func Solve(input string) int {
	input = input + string('\n')
	for _, r := range regexes {
		input = r.ReplaceAllString(input, "")
	}
	ids := gamer.FindAllString(input, -1)

	sum := 0
	for _, id := range ids {
		i, _ := strconv.Atoi(strings.Replace(id, "Game ", "", 1))
		sum += i
	}

	return sum
}
