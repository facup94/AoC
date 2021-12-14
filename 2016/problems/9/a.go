package day9

import (
	"strconv"
)

func A(input []string) (interface{}, error) {
	chars := []rune(input[0])
	var decompressed []rune

	var i int
	for i < len(chars) {
		if chars[i] != '(' {
			decompressed = append(decompressed, chars[i])
			i++
			continue
		}

		i++
		x := i
		for chars[x] != 'x' {
			x++
		}
		size, _ := strconv.Atoi(string(chars[i:x]))

		x++
		i = x
		for chars[x] != ')' {
			x++
		}
		repetitions, _ := strconv.Atoi(string(chars[i:x]))

		i = x + 1
		for n := 0; n < repetitions; n++ {
			decompressed = append(decompressed, chars[i:i+size]...)
		}
		i += size
	}

	return len(decompressed), nil
}
