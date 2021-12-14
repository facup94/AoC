package day10

import (
	"errors"
)

const chipToCompareA, chipToCompareB int = 61, 17

func A(input []string) (interface{}, error) {
	bots, instructions := getBotsAndInstructionsFromInput(input)

	var i int
	for i < len(instructions) {
		chip, id := instructions[i].chip, instructions[i].botID

		bots[id] = bots[id].putChip(chip)
		if bots[id].hasBothChips() {
			if bots[id].chipsAre(chipToCompareA, chipToCompareB) {
				return id, nil
			}

			tempBot, low, high := bots[id].removeChips()
			bots[id] = tempBot

			if bots[id].lowRule.destination == "bot" {
				instructions = append(instructions, putChipInBotInstruction{chip: low, botID: bots[id].lowRule.n})
			}
			if bots[id].highRule.destination == "bot" {
				instructions = append(instructions, putChipInBotInstruction{chip: high, botID: bots[id].highRule.n})
			}
		}
		i++
	}

	return nil, errors.New("solution not found")
}
