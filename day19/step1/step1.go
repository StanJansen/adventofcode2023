package step1

import (
	"strconv"
	"strings"
)

type Workflows map[Action]Workflow

type Workflow []Instruction

type Instruction struct {
	Condition *Condition
	Action    Action
}

type Condition struct {
	Type          Type
	ConditionType ConditionType
	Amount        Amount
}

type Type rune

type Action string

type ConditionType rune

type Amount uint16

type Ratings map[Type]Amount

const (
	START_ACTION    Action = "in"
	ACCEPTED_ACTION Action = "A"
	REJECTED_ACTION Action = "R"

	EXTREME_TYPE     Type = 'x'
	MUSICAL_TYPE     Type = 'm'
	AERODYNAMIC_TYPE Type = 'a'
	SHINY_TYPE       Type = 's'

	GREATER_THAN ConditionType = '>'
	LESS_THAN    ConditionType = '<'
)

func Solve(input string) int {
	lines := strings.Split(input, "\n")
	workflows, li := createWorkflows(lines)

	sum := 0
	for _, line := range lines[li:] {
		ratings := Ratings{}
		line = strings.TrimLeft(line, "{")
		line = strings.TrimRight(line, "}")
		for _, part := range strings.Split(line, ",") {
			p := strings.Split(part, "=")
			a, _ := strconv.Atoi(p[1])
			ratings[Type(p[0][0])] = Amount(a)
		}

		a := START_ACTION
		for a != ACCEPTED_ACTION && a != REJECTED_ACTION {
			a = workflows[a].Execute(ratings)
		}

		if a == ACCEPTED_ACTION {
			for _, r := range ratings {
				sum += int(r)
			}
		}
	}

	return sum
}

func createWorkflows(lines []string) (workflows Workflows, li int) {
	workflows = Workflows{}
	for _, line := range lines {
		li++
		if line == "" {
			break
		}

		parts := strings.Split(line, "{")
		action := Action(parts[0])
		for _, part := range strings.Split(strings.TrimRight(parts[1], "}"), ",") {
			p := strings.Split(part, ":")
			if len(p) == 1 {
				workflows[action] = append(workflows[action], Instruction{
					Action: Action(p[0]),
				})
				continue
			}

			a, _ := strconv.Atoi(p[0][2:])
			workflows[action] = append(workflows[action], Instruction{
				Condition: &Condition{
					Type:          Type(p[0][0]),
					ConditionType: ConditionType(p[0][1]),
					Amount:        Amount(a),
				},
				Action: Action(p[1]),
			})
		}
	}

	return
}

func (w Workflow) Execute(ratings Ratings) Action {
	for _, i := range w {
		if i.Condition != nil && !i.Condition.Compare(ratings) {
			continue
		}

		return i.Action
	}

	return ACCEPTED_ACTION
}

func (c Condition) Compare(ratings Ratings) bool {
	switch c.ConditionType {
	case GREATER_THAN:
		return ratings[c.Type] > c.Amount
	default:
		return ratings[c.Type] < c.Amount
	}
}
