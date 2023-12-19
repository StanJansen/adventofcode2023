package step2

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

type RatingMap map[Type]Rating

type Rating struct {
	Min, Max Amount
}

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
	workflows := createWorkflows(lines)
	ratingMaps := workflows.CreateRatingMaps(START_ACTION, RatingMap{
		EXTREME_TYPE:     {Min: 1, Max: 4000},
		MUSICAL_TYPE:     {Min: 1, Max: 4000},
		AERODYNAMIC_TYPE: {Min: 1, Max: 4000},
		SHINY_TYPE:       {Min: 1, Max: 4000},
	})

	sum := 0
	for _, rm := range ratingMaps {
		rmSum := 0
		for _, rating := range rm {
			if rmSum == 0 {
				rmSum = int(rating.Max - rating.Min + 1)
			} else {
				rmSum *= int(rating.Max - rating.Min + 1)
			}
		}
		sum += rmSum
	}
	return sum
}

func createWorkflows(lines []string) Workflows {
	workflows := Workflows{}
	for _, line := range lines {
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

	return workflows
}

func (w Workflows) CreateRatingMaps(action Action, ratingMap RatingMap) []RatingMap {
	ratingMaps := []RatingMap{ratingMap}
	result := []RatingMap{}
	for _, i := range w[action] {
		if i.Condition == nil {
			if i.Action == ACCEPTED_ACTION {
				result = append(result, ratingMaps...)
			} else {
				for _, ratingMap := range ratingMaps {
					result = append(result, w.CreateRatingMaps(i.Action, ratingMap)...)
				}
			}

			break
		}

		rm := ratingMaps
		ratingMaps = []RatingMap{}
		for _, ratingMap := range rm {
			successRm, failureRm, successOk, failureOk := i.Condition.SplitRatingMap(ratingMap)
			if successOk {
				if i.Action == ACCEPTED_ACTION {
					result = append(result, successRm)
				} else if i.Action != REJECTED_ACTION {
					result = append(result, w.CreateRatingMaps(i.Action, successRm)...)
				}
			}
			if failureOk {
				ratingMaps = append(ratingMaps, failureRm)
			}
		}
	}
	return result
}

func (c Condition) SplitRatingMap(rm RatingMap) (successRm, failureRm RatingMap, successOk, failureOk bool) {
	successRm = RatingMap{}
	failureRm = RatingMap{}
	for t, rating := range rm {
		successRm[t] = rating
		failureRm[t] = rating
	}

	min := rm[c.Type].Min
	max := rm[c.Type].Max

	switch c.ConditionType {
	case GREATER_THAN:
		if max <= c.Amount {
			failureOk = true
		} else if min > c.Amount {
			successOk = true
		} else {
			successOk = true
			failureOk = true
			successRm[c.Type] = Rating{
				Min: c.Amount + 1,
				Max: max,
			}
			failureRm[c.Type] = Rating{
				Min: min,
				Max: c.Amount,
			}
		}
	case LESS_THAN:
		if min >= c.Amount {
			failureOk = true
		} else if max < c.Amount {
			successOk = true
		} else {
			successOk = true
			failureOk = true
			successRm[c.Type] = Rating{
				Min: min,
				Max: c.Amount - 1,
			}
			failureRm[c.Type] = Rating{
				Min: c.Amount,
				Max: max,
			}
		}
	}

	return
}
