package day10

import (
	"fmt"
	"strconv"
	"strings"

	"AoC/commons"
)

func getBotsAndInstructionsFromInput(input []string) (bots []robot, instructions []putChipInBotInstruction) {
	for _, s := range input {
		if strings.HasPrefix(s, "value") {
			fields := strings.Fields(s)
			chip, _ := strconv.Atoi(fields[1])
			id, _ := strconv.Atoi(fields[5])

			instructions = append(instructions, putChipInBotInstruction{chip: chip, botID: id})
		} else if strings.HasPrefix(s, "bot") {
			fields := strings.Fields(s)

			botID, _ := strconv.Atoi(fields[1])
			lowDst := fields[5]
			lowDstN, _ := strconv.Atoi(fields[6])
			highDst := fields[10]
			highDstN, _ := strconv.Atoi(fields[11])

			for len(bots) <= botID {
				bots = append(bots, robot{id: len(bots), low: -1, high: -1})
			}

			bots[botID] = robot{
				id:       botID,
				low:      -1,
				high:     -1,
				lowRule:  rule{destination: lowDst, n: lowDstN},
				highRule: rule{destination: highDst, n: highDstN},
			}
		}
	}

	return bots, instructions
}

type robot struct {
	id       int
	low      int
	high     int
	lowRule  rule
	highRule rule
}

type rule struct {
	destination string
	n           int
}

func (r robot) String() string {
	return fmt.Sprintf("{[%d] - low: %d -> %s - high: %d -> %s}", r.id, r.low, r.lowRule, r.high, r.highRule)
}

func (r robot) putChip(chip int) robot {
	if r.low == -1 && r.high == -1 {
		r.low = chip
		return r
	}

	if r.low != -1 {
		r.low, r.high = commons.IntMin(r.low, chip), commons.IntMax(r.low, chip)
	} else if r.high != -1 {
		r.low, r.high = commons.IntMin(r.high, chip), commons.IntMax(r.high, chip)
	}
	return r
}

func (r robot) hasBothChips() bool {
	return r.low != -1 && r.high != -1
}

func (r robot) chipsAre(a int, b int) bool {
	return r.low == commons.IntMin(a, b) && r.high == commons.IntMax(a, b)
}

func (r robot) removeChips() (bot robot, low int, high int) {
	low, high = r.low, r.high
	r.low, r.high = -1, -1
	return r, low, high
}

func (r rule) String() string {
	return fmt.Sprintf("{%s %d}", r.destination, r.n)
}

type putChipInBotInstruction struct {
	chip  int
	botID int
}
