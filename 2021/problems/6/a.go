package day06

func A(input []string) (interface{}, error) {
	fishes := getFishesFromInput(input[0])

	for day := 0; day < 80; day++ {
		var newfishes []int
		for i := 0; i < len(fishes); i++ {
			if fishes[i] == 0 {
				fishes[i] = 6
				newfishes = append(newfishes, 8)
			} else {
				fishes[i]--
			}
		}
		fishes = append(fishes, newfishes...)
	}

	return len(fishes), nil
}
