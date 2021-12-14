package day1

import (
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	rotation string
	blocks   int
}

func getInstructionList(in string) (instructions []instruction) {
	for _, s := range strings.Split(in, ", ") {
		n, _ := strconv.Atoi(s[1:])
		instructions = append(instructions, instruction{rotation: s[0:1], blocks: n})
	}
	return instructions
}

type direction string

const (
	N direction = "N"
	E direction = "E"
	S direction = "S"
	W direction = "W"
)

func (d direction) Rotate90Clockwise() direction {
	switch d {
	case N:
		return E
	case E:
		return S
	case S:
		return W
	case W:
		return N
	}
	return d
}

func (d direction) Rotate90CounterClockwise() direction {
	switch d {
	case N:
		return W
	case E:
		return N
	case S:
		return E
	case W:
		return S
	}
	return d
}

type position struct {
	x   int
	y   int
	dir direction
}

func (pos *position) followInstruction(i instruction) {
	if i.rotation == "R" {
		pos.dir = pos.dir.Rotate90Clockwise()
	} else if i.rotation == "L" {
		pos.dir = pos.dir.Rotate90CounterClockwise()
	}

	switch pos.dir {
	case N:
		pos.y += i.blocks
	case E:
		pos.x += i.blocks
	case S:
		pos.y -= i.blocks
	case W:
		pos.x -= i.blocks
	}
}

func hashFromCoordinates(x, y int) string {
	return fmt.Sprintf("X:%d|Y:%d", x, y)
}
