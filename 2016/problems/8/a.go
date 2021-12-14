package day8

func A(input []string) (interface{}, error) {
	var scrn screen
	for _, s := range input {
		instr := stringToInstruction(s)
		scrn = instr(scrn)
	}

	var countOn int
	for _, row := range scrn {
		for _, pixel := range row {
			if pixel {
				countOn++
			}
		}
	}

	return countOn, nil
}
