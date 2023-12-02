package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var rr = regexp.MustCompile("(\\d{1,})( red)")
var gr = regexp.MustCompile("(\\d{1,})( green)")
var br = regexp.MustCompile("(\\d{1,})( blue)")

func main() {
	input := os.Args[1]
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		r := rr.FindAllString(line, -1)
		g := gr.FindAllString(line, -1)
		b := br.FindAllString(line, -1)

		lr := -1
		lg := -1
		lb := -1

		for _, v := range r {
			v = strings.Replace(v, " red", "", -1)
			amount, _ := strconv.Atoi(v)
			if lr == -1 || amount > lr {
				lr = amount
			}
		}

		for _, v := range g {
			v = strings.Replace(v, " green", "", -1)
			amount, _ := strconv.Atoi(v)
			if lg == -1 || amount > lg {
				lg = amount
			}
		}

		for _, v := range b {
			v = strings.Replace(v, " blue", "", -1)
			amount, _ := strconv.Atoi(v)
			if lb == -1 || amount > lb {
				lb = amount
			}
		}

		sum += lr * lg * lb
	}

	fmt.Println(sum)
}
