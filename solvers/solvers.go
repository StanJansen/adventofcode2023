package solvers

import (
	d1s1 "github.com/StanJansen/adventofcode2023/day1/step1"
	d1s2 "github.com/StanJansen/adventofcode2023/day1/step2"
	d10s1 "github.com/StanJansen/adventofcode2023/day10/step1"
	d10s2 "github.com/StanJansen/adventofcode2023/day10/step2"
	d11s1 "github.com/StanJansen/adventofcode2023/day11/step1"
	d11s2 "github.com/StanJansen/adventofcode2023/day11/step2"
	d2s1 "github.com/StanJansen/adventofcode2023/day2/step1"
	d2s2 "github.com/StanJansen/adventofcode2023/day2/step2"
	d3s1 "github.com/StanJansen/adventofcode2023/day3/step1"
	d3s2 "github.com/StanJansen/adventofcode2023/day3/step2"
	d4s1 "github.com/StanJansen/adventofcode2023/day4/step1"
	d4s2 "github.com/StanJansen/adventofcode2023/day4/step2"
	d5s1 "github.com/StanJansen/adventofcode2023/day5/step1"
	d5s2 "github.com/StanJansen/adventofcode2023/day5/step2"
	d6s1 "github.com/StanJansen/adventofcode2023/day6/step1"
	d6s2 "github.com/StanJansen/adventofcode2023/day6/step2"
	d7s1 "github.com/StanJansen/adventofcode2023/day7/step1"
	d7s2 "github.com/StanJansen/adventofcode2023/day7/step2"
	d8s1 "github.com/StanJansen/adventofcode2023/day8/step1"
	d8s2 "github.com/StanJansen/adventofcode2023/day8/step2"
	d9s1 "github.com/StanJansen/adventofcode2023/day9/step1"
	d9s2 "github.com/StanJansen/adventofcode2023/day9/step2"
)

var Solvers = [][2]func(input string) int{
	{d1s1.Solve, d1s2.Solve},
	{d2s1.Solve, d2s2.Solve},
	{d3s1.Solve, d3s2.Solve},
	{d4s1.Solve, d4s2.Solve},
	{d5s1.Solve, d5s2.Solve},
	{d6s1.Solve, d6s2.Solve},
	{d7s1.Solve, d7s2.Solve},
	{d8s1.Solve, d8s2.Solve},
	{d9s1.Solve, d9s2.Solve},
	{d10s1.Solve, d10s2.Solve},
	{d11s1.Solve, d11s2.Solve},
}
