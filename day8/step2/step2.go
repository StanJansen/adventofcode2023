package step2

import "strings"

func Solve(input string) int {
	lines := strings.Split(input, "\n")

	m := make(map[string][2]string, len(lines)-2)
	nodes := []string{}
	for _, line := range lines[2:] {
		node := line[0:3]
		m[node] = [2]string{line[7:10], line[12:15]}
		if line[2] == 'A' {
			nodes = append(nodes, node)
		}
	}

	counts := []int{}
	for _, n := range nodes {
		var s int
		for n != n[0:2]+"Z" {
			for _, i := range lines[0] {
				s++
				if i == 'L' {
					n = m[n][0]
				} else {
					n = m[n][1]
				}
			}
		}
		counts = append(counts, s)
	}

	return getLeastCommonMultiple(counts)
}

func getLeastCommonMultiple(numbers []int) int {
	lcm := numbers[0]
	for i := 0; i < len(numbers); i++ {
		num1 := lcm
		num2 := numbers[i]
		gcd := 1
		for num2 != 0 {
			temp := num2
			num2 = num1 % num2
			num1 = temp
		}
		gcd = num1
		lcm = (lcm * numbers[i]) / gcd
	}

	return lcm
}
