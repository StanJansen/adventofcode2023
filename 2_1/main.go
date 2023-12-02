package two_one

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var rr = regexp.MustCompile("(Game \\d{1,}).*([2-9][0-9]|1[3-9]|\\d{3,})( red).*\n")
var gr = regexp.MustCompile("(Game \\d{1,}).*([2-9][0-9]|1[4-9]|\\d{3,})( green).*\n")
var br = regexp.MustCompile("(Game \\d{1,}).*([2-9][0-9]|1[5-9]|\\d{3,})( blue).*\n")
var gamer = regexp.MustCompile("Game (\\d{1,})")

func main() {
	input := os.Args[1] + "\n"
	input = rr.ReplaceAllString(input, "")
	input = gr.ReplaceAllString(input, "")
	input = br.ReplaceAllString(input, "")

	ids := gamer.FindAllString(input, -1)

	sum := 0

	for _, id := range ids {
		i, _ := strconv.Atoi(strings.Replace(id, "Game ", "", 1))
		sum += i
	}

	fmt.Println(sum)
}
