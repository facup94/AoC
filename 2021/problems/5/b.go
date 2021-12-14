package day05

import (
	"AoC/commons"
)

func B(input []string) (interface{}, error) {
	lines := inputToLines(input)

	diagram := make(map[point]int)

	for _, l := range lines {
		if l.from.x != l.to.x && l.from.y != l.to.y {
			// force line to go left -> right
			if l.from.x > l.to.x {
				l.from, l.to = l.to, l.from
			}

			yDirection := -1 // Example: \
			if l.from.y < l.to.y {
				yDirection = 1 // Example: /
			}

			for x := l.from.x; x < l.to.x+1; x++ {
				y := l.from.y + ((x - l.from.x) * yDirection)
				diagram[point{x, y}]++
			}
		}

		if l.from.y == l.to.y {
			for x := commons.IntMin(l.from.x, l.to.x); x < commons.IntMax(l.from.x, l.to.x)+1; x++ {
				diagram[point{x: x, y: l.from.y}]++
			}
		}

		if l.from.x == l.to.x {
			for y := commons.IntMin(l.from.y, l.to.y); y < commons.IntMax(l.from.y, l.to.y)+1; y++ {
				diagram[point{x: l.from.x, y: y}]++
			}
		}
	}

	var count int
	for _, v := range diagram {
		if v >= 2 {
			count++
		}
	}

	return count, nil
}
