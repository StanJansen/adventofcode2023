package main

import (
	"fmt"
	"os"

	"github.com/StanJansen/adventofcode2023/day1/step1"
	"github.com/StanJansen/adventofcode2023/day1/step2"
)

func main() {
	content, _ := os.ReadFile("./day1/input.txt")
	input := string(content)

	fmt.Printf("Day 1 step 1: %d\n", step1.Solve(input))
	fmt.Printf("Day 1 step 2: %d\n", step2.Solve(input))
}
