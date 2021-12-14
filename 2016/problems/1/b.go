package day1

import (
	"errors"

	"AoC/commons"
)

func B(input []string) (interface{}, error) {
	history := make(map[string]bool)

	p := position{x: 0, y: 0, dir: N}
	history[hashFromCoordinates(0, 0)] = true

	for _, i := range getInstructionList(input[0]) {
		fromX, fromY := p.x, p.y
		p.followInstruction(i)
		toX, toY := p.x, p.y

		if fromX != toX { // horizontal movement
			if fromX < toX { // eastbound movement
				for x := fromX + 1; x < toX+1; x++ {
					if visited := history[hashFromCoordinates(x, fromY)]; visited {
						return commons.IntAbs(x) + commons.IntAbs(fromY), nil
					}
					history[hashFromCoordinates(x, fromY)] = true
				}
			} else { // westbound movement
				for x := fromX - 1; x >= toX; x-- {
					if visited := history[hashFromCoordinates(x, fromY)]; visited {
						return commons.IntAbs(x) + commons.IntAbs(fromY), nil
					}
					history[hashFromCoordinates(x, fromY)] = true
				}
			}
		}

		if fromY != toY { // vertical movement
			if fromY < toY { // northbound movement
				for y := fromY + 1; y < toY+1; y++ {
					if visited := history[hashFromCoordinates(fromX, y)]; visited {
						return commons.IntAbs(fromX) + commons.IntAbs(y), nil
					}
					history[hashFromCoordinates(fromX, y)] = true
				}
			} else { // southbound movement
				for y := fromY - 1; y >= toY; y-- {
					if visited := history[hashFromCoordinates(fromX, y)]; visited {
						return commons.IntAbs(fromX) + commons.IntAbs(y), nil
					}
					history[hashFromCoordinates(fromX, y)] = true
				}
			}
		}
	}

	return -1, errors.New("no position visited twice")
}
