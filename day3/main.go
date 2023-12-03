package main

import (
	"fmt"
	"os"

	"github.com/StanJansen/adventofcode2023/day3/step1"
	"github.com/StanJansen/adventofcode2023/day3/step2"
)

func main() {
	content, _ := os.ReadFile("./day3/input.txt")
	input := string(content)

	fmt.Printf("Day 3 step 1: %d\n", step1.Solve(input))
	fmt.Printf("Day 3 step 2: %d\n", step2.Solve(input))
}
