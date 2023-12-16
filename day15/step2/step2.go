package step2

import (
	"slices"
	"strconv"
	"strings"
)

type Lens struct {
	key   string
	value int
}

func Solve(input string) int {
	steps := strings.Split(input, ",")
	boxes := [256][]Lens{}
	for _, step := range steps {
		if step[len(step)-1] == '-' {
			key := step[:len(step)-1]
			box := getBox(key)
			index := slices.IndexFunc(boxes[box], func(l Lens) bool { return l.key == key })
			if index != -1 {
				boxes[box] = append(boxes[box][:index], boxes[box][index+1:]...)
			}
		} else {
			key := step[:len(step)-2]
			value, _ := strconv.Atoi(string(step[len(step)-1]))
			box := getBox(key)
			index := slices.IndexFunc(boxes[box], func(l Lens) bool { return l.key == key })
			if index == -1 {
				boxes[box] = append(boxes[box], Lens{key, value})
			} else {
				boxes[box][index].value = value
			}
		}
	}

	sum := 0
	for box, lenses := range boxes {
		for idx, lens := range lenses {
			sum += (box + 1) * (idx + 1) * lens.value
		}
	}
	return sum
}

func getBox(key string) (box int) {
	for _, char := range key {
		box = (box + int(char)) * 17 % 256
	}
	return
}
