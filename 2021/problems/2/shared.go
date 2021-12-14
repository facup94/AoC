package day02

import (
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	units     int
}

func parseInput(input []string) ([]instruction, error) {
	var instructions []instruction
	for _, s := range input {
		ss := strings.Split(s, " ")
		u, err := strconv.Atoi(ss[1])
		if err != nil {
			return nil, err
		}

		instructions = append(instructions, instruction{direction: ss[0], units: u})
	}

	return instructions, nil
}
