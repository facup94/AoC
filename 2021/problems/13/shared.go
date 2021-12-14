package day13

import (
	"fmt"
	"strconv"
	"strings"
)

type dot struct {
	x int
	y int
}

func getDotsFromInput(input []string) (dots []dot) {
	var i int
	for {
		if input[i] == "" {
			break
		}

		ss := strings.Split(input[i], ",")
		x, _ := strconv.Atoi(ss[0])
		y, _ := strconv.Atoi(ss[1])
		dots = append(dots, dot{x: x, y: y})
		i++
	}

	return dots
}

type fold struct {
	direction string
	position  int
}

func getFoldsFromInput(input []string) (folds []fold) {
	var i int
	for input[i] != "" {
		i++
	}

	for _, s := range input[i+1:] {
		ss := strings.Split(strings.Fields(s)[2], "=")
		pos, _ := strconv.Atoi(ss[1])
		folds = append(folds, fold{direction: ss[0], position: pos})
	}

	return folds
}

func prettyPrint(dots []dot) string {
	var maxX, maxY int
	dotsMap := make(map[string]bool)

	for _, d := range dots {
		if d.x > maxX {
			maxX = d.x
		}
		if d.y > maxY {
			maxY = d.y
		}

		dotsMap[fmt.Sprintf("%d|%d", d.x, d.y)] = true
	}

	var sb strings.Builder
	for y := 0; y < maxY+1; y++ {
		for x := 0; x < maxX+1; x++ {
			if dotsMap[fmt.Sprintf("%d|%d", x, y)] {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}
