package step2

import (
	"slices"
	"strconv"
	"strings"
)

type cardResult struct {
	gameNr int
	sum    int
}

var cards = []*cardResult{}
var sums = map[int]int{}
var sum = 0

func Solve(input string) int {
	lines := strings.Split(strings.ReplaceAll(input, "Card ", ""), "\n")

	for _, line := range lines {
		cardParts := strings.Split(line, ":")
		gameNr, _ := strconv.Atoi(strings.TrimLeft(cardParts[0], " "))
		card := cardResult{gameNr, 0}
		cards = append(cards, &card)
		solveCard(&card, cardParts[1])
	}

	solveCards(cards)

	return sum
}

func solveCards(c []*cardResult) {
	sum += len(c)
	for _, card := range c {
		if card.sum > 0 {
			solveCards(cards[card.gameNr : card.gameNr+card.sum])
		}
	}
}

func solveCard(card *cardResult, input string) {
	parts := strings.Split(input, "|")

	wNumbers := parseDigits(parts[0])
	cNumbers := parseDigits(parts[1])

	for _, n := range cNumbers {
		if slices.Contains(wNumbers, n) {
			card.sum++
		}
	}
}

func parseDigits(input string) []int {
	digits := []int{}
	for _, char := range strings.Split(input, " ") {
		char = strings.Trim(char, " ")
		if char == "" {
			continue
		}
		digit, _ := strconv.Atoi(char)
		digits = append(digits, digit)
	}
	return digits
}
