package step1

import (
	"strings"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

type Component struct {
	Name        string
	Connections map[string]*Component
}

type Node struct {
	Component *Component
	Distance  int
}

// You might need to run this a few times, it's the lowest amount (this input gives both 2259009 and 562772 for example as results).
func Solve(input string) int {
	components := getComponents(input)
	distances := make(map[string]int)

	for _, component := range components {
		for _, connection := range component.Connections {
			if _, ok := distances[connection.Name+component.Name]; ok {
				continue
			}
			distances[component.Name+connection.Name] = getDistance(component, connection, components)
		}
	}

	highestThree := make([]string, 3)
	highestThreeDistances := make([]int, 3)
	for name, distance := range distances {
		for i, highestDistance := range highestThreeDistances {
			if distance > highestDistance {
				highestThree[i] = name
				highestThreeDistances[i] = distance
				break
			}
		}
	}

	for _, name := range highestThree {
		name1 := name[0:3]
		name2 := name[3:6]
		delete(components[name1].Connections, name2)
		delete(components[name2].Connections, name1)
	}

	group1Start := highestThree[0][0:3]
	group2Start := highestThree[0][3:6]

	return getGroupLength(group1Start, components) * getGroupLength(group2Start, components)
}

func getComponents(input string) map[string]*Component {
	lines := strings.Split(input, "\n")
	components := make(map[string]*Component)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		name := parts[0]
		component, ok := components[name]
		if !ok {
			component = &Component{Name: name, Connections: make(map[string]*Component)}
			components[name] = component
		}

		connections := strings.Split(parts[1], " ")
		for _, connection := range connections {
			conn, ok := components[connection]
			if !ok {
				conn = &Component{Name: connection, Connections: make(map[string]*Component)}
				components[connection] = conn
			}
			component.Connections[connection] = conn
			conn.Connections[name] = component
		}
	}
	return components
}

func getDistance(from, to *Component, components map[string]*Component) int {
	distances := make(map[*Component]int)
	visited := make(map[*Component]bool)
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		return a.(*Node).Distance - b.(*Node).Distance
	})

	for _, component := range components {
		distances[component] = int(^uint(0) >> 1)
	}
	distances[from] = 0

	pq.Enqueue(&Node{Component: from, Distance: 0})

	for !pq.Empty() {
		v, _ := pq.Dequeue()
		node := v.(*Node)
		current := node.Component

		if visited[current] {
			continue
		}
		visited[current] = true

		for _, connection := range current.Connections {
			if current == from && connection == to {
				continue
			}
			distance := distances[current] + 1
			if distance < distances[connection] {
				distances[connection] = distance
				pq.Enqueue(&Node{Component: connection, Distance: distance})
			}
		}
	}

	return distances[to]
}

func getGroupLength(start string, components map[string]*Component) int {
	visited := make(map[*Component]bool)
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		return a.(*Node).Distance - b.(*Node).Distance
	})

	pq.Enqueue(&Node{Component: components[start], Distance: 0})

	for !pq.Empty() {
		v, _ := pq.Dequeue()
		node := v.(*Node)
		current := node.Component

		if visited[current] {
			continue
		}
		visited[current] = true

		for _, connection := range current.Connections {
			pq.Enqueue(&Node{Component: connection, Distance: node.Distance + 1})
		}
	}

	return len(visited)
}
