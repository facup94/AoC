package day10

func B(input []string) (interface{}, error) {
	outputs := make(map[int]int)

	bots, instructions := getBotsAndInstructionsFromInput(input)

	var i int
	for i < len(instructions) {
		chip, id := instructions[i].chip, instructions[i].botID

		bots[id] = bots[id].putChip(chip)
		if bots[id].hasBothChips() {
			tempBot, low, high := bots[id].removeChips()
			bots[id] = tempBot

			if bots[id].lowRule.destination == "bot" {
				instructions = append(instructions, putChipInBotInstruction{chip: low, botID: bots[id].lowRule.n})
			} else {
				outputs[bots[id].lowRule.n] = low
			}
			if bots[id].highRule.destination == "bot" {
				instructions = append(instructions, putChipInBotInstruction{chip: high, botID: bots[id].highRule.n})
			} else {
				outputs[bots[id].highRule.n] = high
			}
		}
		i++
	}

	return outputs[0] * outputs[1] * outputs[2], nil
}
