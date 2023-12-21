package step1

import (
	"strings"
)

type Modules map[string]*Module
type Module struct {
	Type         ModuleType
	State        bool
	Destinations []string
	Inputs       map[string]Pulse
}
type Signal struct {
	From  string
	To    string
	Pulse Pulse
}
type ModuleType rune
type Pulse byte
type Queue []Signal

const (
	BROADCASTER = ModuleType('b')
	FLIPFLOP    = ModuleType('%')
	CONJUCTION  = ModuleType('&')

	LOW  = Pulse(0)
	HIGH = Pulse(1)
)

func Solve(input string) int {
	modules := parseInput(input)

	low, high := 0, 0
	for i := 0; i < 1000; i++ {
		l, h := modules.PressButton()
		low += l
		high += h
	}

	return low * high
}

func parseInput(input string) Modules {
	input = strings.ReplaceAll(input, " ", "")
	lines := strings.Split(input, "\n")

	modules := make(Modules, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, "->")
		destinations := strings.Split(parts[1], ",")
		module := Module{
			Type:         ModuleType(parts[0][0]),
			Destinations: destinations,
			Inputs:       map[string]Pulse{},
		}
		key := parts[0][1:]
		if module.Type == BROADCASTER {
			key = parts[0]
		}
		modules[key] = &module
	}

	for key, module := range modules {
		for _, destination := range module.Destinations {
			if _, ok := modules[destination]; ok {
				modules[destination].Inputs[key] = LOW
			}
		}
	}

	return modules
}

func (modules Modules) PressButton() (int, int) {
	queue := Queue{}
	for _, d := range modules["broadcaster"].Destinations {
		queue = append(queue, Signal{
			From:  "broadcaster",
			To:    d,
			Pulse: LOW,
		})
	}

	return queue.Read(modules, 1, 0)
}

func (queue Queue) Read(modules Modules, l, h int) (int, int) {
	newQueue := Queue{}
	for _, signal := range queue {
		if signal.Pulse == LOW {
			l++
		} else {
			h++
		}

		to, ok := modules[signal.To]
		if !ok {
			continue
		}

		hasPulse := false
		pulse := LOW
		if to.Type == FLIPFLOP && signal.Pulse == LOW {
			to.State = !to.State
			hasPulse = true
			if to.State {
				pulse = HIGH
			}
		} else if to.Type == CONJUCTION {
			to.Inputs[signal.From] = signal.Pulse
			hasPulse = true
			for _, input := range to.Inputs {
				if input == LOW {
					pulse = HIGH
					break
				}
			}
		}

		if hasPulse {
			for _, destination := range to.Destinations {
				newQueue = append(newQueue, Signal{
					From:  signal.To,
					To:    destination,
					Pulse: pulse,
				})
			}
		}
	}

	if len(newQueue) == 0 {
		return l, h
	}

	return newQueue.Read(modules, l, h)
}
