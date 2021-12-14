package day06

func B(input []string) (interface{}, error) {
	fishes := getFishesFromInput(input[0])

	// numInternalTimer[i] is the number of fishes with internal timer = i
	var numInternalTimer [9]int
	for _, n := range fishes {
		numInternalTimer[n]++
	}

	for day := 0; day < 256; day++ {
		numInternalTimer[0], numInternalTimer[1], numInternalTimer[2], numInternalTimer[3], numInternalTimer[4], numInternalTimer[5], numInternalTimer[6], numInternalTimer[7], numInternalTimer[8] =
			numInternalTimer[1], numInternalTimer[2], numInternalTimer[3], numInternalTimer[4], numInternalTimer[5], numInternalTimer[6], numInternalTimer[7]+numInternalTimer[0], numInternalTimer[8], numInternalTimer[0]
	}

	var sum int
	for _, n := range numInternalTimer {
		sum += n
	}
	return sum, nil
}
