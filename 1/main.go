package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var chars = "zero|one|two|three|four|five|six|seven|eight|nine"
var regex = regexp.MustCompile("[0-9]|" + chars)
var reverseRegex = regexp.MustCompile("[0-9]|" + reverse(chars))
var mapping = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	input := os.Args[1]
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		first := regex.FindString(line)
		last := reverse(reverseRegex.FindString(reverse(line)))
		if val, ok := mapping[first]; ok {
			first = val
		}
		if val, ok := mapping[last]; ok {
			last = val
		}

		amount, _ := strconv.Atoi(first + last)

		sum += amount
	}

	fmt.Println(sum)
}

func reverse(value string) string {
	data := []rune(value)
	result := []rune{}

	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}

	return string(result)
}
