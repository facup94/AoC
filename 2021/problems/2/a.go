package day02

func A(input []string) (interface{}, error) {
	instructions, err := parseInput(input)
	if err != nil {
		return -1, err
	}

	var h, depth int
	for _, i := range instructions {
		switch i.direction {
		case "forward":
			h += i.units
		case "down":
			depth += i.units
		case "up":
			depth -= i.units
		}
	}

	return h * depth, nil
}
