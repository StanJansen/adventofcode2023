package step2

import (
	"strconv"
	"strings"
)

const (
	UNSOLVED     = -1
	UNFOLD_COUNT = 5
)

type SolutionInput struct {
	Chars     []rune
	Groups    []byte
	CharIdx   byte
	GroupIdx  byte
	Solutions [][]int
}

func (input *SolutionInput) withIndexes(charIdx, groupIdx byte) *SolutionInput {
	return &SolutionInput{Chars: input.Chars, Groups: input.Groups, CharIdx: charIdx, GroupIdx: groupIdx, Solutions: input.Solutions}
}

func (input *SolutionInput) getSolution() int {
	return input.Solutions[input.CharIdx][input.GroupIdx]
}

func (input *SolutionInput) storeSolution(solution int) {
	input.Solutions[input.CharIdx][input.GroupIdx] = solution
}

func Solve(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for idx, line := range lines {
		sum += solveLine(idx, line)
	}

	return sum
}

func solveLine(idx int, line string) int {
	parts := strings.Split(line, " ")
	input := &SolutionInput{Chars: []rune(parts[0])}

	for _, group := range strings.Split(parts[1], ",") {
		v, _ := strconv.Atoi(group)
		input.Groups = append(input.Groups, byte(v))
	}

	input.Solutions = make([][]int, len(input.Chars))
	for cIdx := range input.Chars {
		input.Solutions[cIdx] = make([]int, len(input.Groups)+1)
		for gIdx := 0; gIdx <= len(input.Groups); gIdx++ {
			input.Solutions[cIdx][gIdx] = UNSOLVED
		}
	}

	return input.getCount()
}

func (input *SolutionInput) getCount() (count int) {
	if input.CharIdx >= byte(len(input.Chars)) {
		if input.GroupIdx >= byte(len(input.Groups)) {
			count++
		}
		return
	}

	if input.getSolution() != UNSOLVED {
		return input.getSolution()
	}

	switch input.Chars[input.CharIdx] {
	case '.':
		return input.withIndexes(input.CharIdx+1, input.GroupIdx).getCount()
	case '?':
		count += input.withIndexes(input.CharIdx+1, input.GroupIdx).getCount()
	}

	if input.GroupIdx >= byte(len(input.Groups)) {
		return
	}

	var matches byte
	for _, char := range input.Chars[input.CharIdx:] {
		if char == '.' || (char == '?' && matches == input.Groups[input.GroupIdx]) {
			break
		}
		matches++
	}
	if matches != input.Groups[input.GroupIdx] {
		return
	}

	nextIdx := input.CharIdx + matches
	if nextIdx < byte(len(input.Chars)) && input.Chars[nextIdx] != '#' {
		nextIdx++
	}

	count += input.withIndexes(nextIdx, input.GroupIdx+1).getCount()
	input.storeSolution(count)

	return
}
