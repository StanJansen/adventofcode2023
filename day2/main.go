package main

import (
	"fmt"
	"os"

	"github.com/StanJansen/adventofcode2023/day2/step1"
	"github.com/StanJansen/adventofcode2023/day2/step2"
)

func main() {
	content, _ := os.ReadFile("./day2/input.txt")
	input := string(content)

	fmt.Printf("Day 2 step 1: %d\n", step1.Solve(input))
	fmt.Printf("Day 2 step 2: %d\n", step2.Solve(input))
}
