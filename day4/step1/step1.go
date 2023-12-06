package step1

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type Solver struct{}

func (Solver) Solve(input string) int {
	cards := strings.Split(input, "\n")

	sum := 0
	for _, card := range cards {
		sum += solveCard(card)
	}

	return sum
}

func solveCard(card string) int {
	card = strings.Split(card, ":")[1]
	parts := strings.Split(card, "|")

	wNumbers := parseDigits(parts[0])
	cNumbers := parseDigits(parts[1])

	power := 0
	for _, n := range cNumbers {
		if slices.Contains(wNumbers, n) {
			power++
		}
	}

	if power <= 1 {
		return power
	}

	return int(math.Pow(2, float64(power-1)))
}

func parseDigits(input string) []int {
	digits := []int{}
	for _, char := range strings.Split(input, " ") {
		char = strings.ReplaceAll(char, " ", "")
		if char == "" {
			continue
		}
		digit, _ := strconv.Atoi(char)
		digits = append(digits, digit)
	}
	return digits
}
