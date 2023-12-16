package step1

import "strings"

func Solve(input string) int {
	steps := strings.Split(input, ",")
	sum := 0
	for _, step := range steps {
		stepSum := 0
		for _, char := range step {
			stepSum = (stepSum + int(char)) * 17 % 256
		}
		sum += stepSum
	}
	return sum
}
