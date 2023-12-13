package step1

import (
	"strings"
)

func Solve(input string) int {
	puzzles := strings.Split(input, "\n\n")

	sum := 0
	for _, puzzle := range puzzles {
		sum += solvePuzzle(puzzle)
	}

	return sum
}

func solvePuzzle(puzzle string) int {
	lines := strings.Split(puzzle, "\n")

	v := solveVertical(lines)
	if v > 0 {
		return v
	}

	return solveHorizontal(lines) * 100
}

func solveVertical(lines []string) int {
	l := lines[0]
	for i := 1; i < len(l); i++ {
		match := true
		count := 0
		for j := 0; j < min(i, len(l)-i); j++ {
			if l[i-j-1] != l[i+j] {
				match = false
				break
			}
			count++
		}
		if !match {
			continue
		}

		for _, line := range lines[1:] {
			cmpPart := []rune(line[i : i+count])
			for j, char := range line[i-count : i] {
				if char != cmpPart[count-j-1] {
					match = false
					break
				}
			}
		}

		if match {
			return i
		}
	}

	return 0
}

func solveHorizontal(lines []string) int {
	for i := 1; i < len(lines); i++ {
		if lines[i-1] != lines[i] {
			continue
		}

		match := true
		for j := 0; j < min(i, len(lines)-i); j++ {
			if lines[i-j-1] != lines[i+j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}

	return 0
}
