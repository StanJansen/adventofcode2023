package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var regex = regexp.MustCompile("[0-9]")

func main() {
	input := os.Args[1]
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		matches := regex.FindAllString(line, -1)
		first := matches[0]
		last := matches[len(matches)-1]

		amount, _ := strconv.Atoi(first + last)

		sum += amount
	}

	fmt.Println(sum)
}
