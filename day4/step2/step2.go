package step2

import (
	"slices"
	"strconv"
	"strings"
)

var cards []string
var sums = map[int]int{}

func Solve(input string) int {
	cards = strings.Split(input, "\n")

	return solveCards(cards)
}

func solveCards(c []string) int {
	sum := len(c)
	for _, card := range c {
		gameNr, cSum := solveCard(card)

		if cSum > 0 {
			sum += solveCards(cards[gameNr : gameNr+cSum])
		}
	}
	return sum
}

func solveCard(card string) (int, int) {
	cardParts := strings.Split(card, ":")
	gameNr, _ := strconv.Atoi(strings.ReplaceAll(strings.Replace(cardParts[0], "Card ", "", 1), " ", ""))
	if sum, ok := sums[gameNr]; ok {
		return gameNr, sum
	}

	card = cardParts[1]
	parts := strings.Split(card, "|")

	wNumbers := parseDigits(parts[0])
	cNumbers := parseDigits(parts[1])

	sum := 0
	for _, n := range cNumbers {
		if slices.Contains(wNumbers, n) {
			sum++
		}
	}

	sums[gameNr] = sum

	return gameNr, sum
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
