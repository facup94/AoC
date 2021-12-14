package day1

import (
	"AoC/commons"
)

func A(input []string) (interface{}, error) {
	p := position{x: 0, y: 0, dir: N}
	for _, i := range getInstructionList(input[0]) {
		p.followInstruction(i)
	}

	return commons.IntAbs(p.x) + commons.IntAbs(p.y), nil
}
