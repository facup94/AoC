package day05

import (
	"strconv"
	"strings"
)

type line struct {
	from point
	to   point
}

type point struct {
	x int
	y int
}

func inputToLines(input []string) (lines []line) {
	for _, s := range input {
		ss := strings.Split(s, " -> ")
		fromS := strings.Split(ss[0], ",")
		toS := strings.Split(ss[1], ",")

		fromX, _ := strconv.Atoi(fromS[0])
		fromY, _ := strconv.Atoi(fromS[1])
		toX, _ := strconv.Atoi(toS[0])
		toY, _ := strconv.Atoi(toS[1])

		lines = append(lines, line{from: point{x: fromX, y: fromY}, to: point{x: toX, y: toY}})
	}
	return lines
}
