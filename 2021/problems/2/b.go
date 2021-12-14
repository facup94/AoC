package day02

func B(input []string) (interface{}, error) {
	instructions, err := parseInput(input)
	if err != nil {
		return -1, err
	}

	var h, depth, aim int
	for _, i := range instructions {
		switch i.direction {
		case "forward":
			h += i.units
			depth += aim * i.units
		case "down":
			aim += i.units
		case "up":
			aim -= i.units
		}
	}

	return h * depth, nil
}
