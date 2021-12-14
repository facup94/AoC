package day6

func B(input []string) (interface{}, error) {
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
		var minRune rune
		var min int
		for r, n := range repetitions[i] {
			if n < min || min == 0 {
				min = n
				minRune = r
			}
		}
		result[i] = minRune
	}

	return string(result), nil
}
