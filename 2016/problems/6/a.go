package day6

func A(input []string) (interface{}, error) {
	messageSize := len(input[0])

	repetitions := make([]map[rune]int, messageSize)
	for i := 0; i < messageSize; i++ {
		repetitions[i] = make(map[rune]int)
	}

	for _, msg := range input {
		for i, r := range msg {
			repetitions[i][r]++
		}
	}

	result := make([]rune, messageSize)
	for i := 0; i < messageSize; i++ {
		var maxRune rune
		var max int
		for r, n := range repetitions[i] {
			if n >= max {
				max = n
				maxRune = r
			}
		}
		result[i] = maxRune
	}

	return string(result), nil
}
