package step1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var symbolIndexes = map[rune]string{'A': "14", 'K': "13", 'Q': "12", 'T': "10", '9': "09", '8': "08", '7': "07", '6': "06", '5': "05", '4': "04", '3': "03", '2': "02", 'J': "01"}

func Solve(input string) int {
	lines := strings.Split(input, "\n")

	hands := make(map[string]int, len(lines))
	for _, line := range lines {
		index, bid := parseLine(line)
		hands[index] = bid
	}

	keys := make([]string, 0, len(hands))
	for k := range hands {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	total := 0
	for i, k := range keys {
		total += (i + 1) * hands[k]
	}

	return total
}

func parseLine(line string) (index string, bid int) {
	parts := strings.Split(line, " ")
	symbols := parts[0]
	bid, _ = strconv.Atoi(parts[1])

	counts := map[rune]byte{}
	jokers := byte(0)
	for _, s := range symbols {
		index += symbolIndexes[s]
		if s == 'J' {
			jokers++
			continue
		}
		counts[s]++
	}

	firstCount, secondCount := getLargestCounts(counts)
	switch firstCount + jokers {
	case 5:
		// Five of a kind
		index = "6" + index
	case 4:
		// Four of a kind
		index = "5" + index
	case 3:
		if secondCount == 2 {
			// Full house
			index = "4" + index
		} else {
			// Three of a kind
			index = "3" + index
		}
	case 2:
		if secondCount == 2 {
			// Two pairs
			index = "2" + index
		} else {
			// One pair
			index = "1" + index
		}
	default:
		// High card
		index = "0" + index
	}

	if symbols == "88JJJ" {
		fmt.Println(index, bid)
	}

	return
}

func getLargestCounts(counts map[rune]byte) (firstCount byte, secondCount byte) {
	for _, count := range counts {
		if count > firstCount || count == firstCount {
			secondCount = firstCount
			firstCount = count
		} else if count > secondCount {
			secondCount = count
		}
	}

	return
}
