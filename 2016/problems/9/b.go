package day9

import (
	"strconv"
	"strings"
)

func B(input []string) (interface{}, error) {
	return decompress([]rune(input[0])), nil
}

func decompress(chars []rune) (size int) {
	var i, count int
	for i < len(chars) {
		if chars[i] != '(' {
			i++
			count++
			continue
		}

		xIdx := i
		for chars[xIdx] != 'x' {
			xIdx++
		}
		numChars, _ := strconv.Atoi(string(chars[i+1 : xIdx]))

		markerCloseIdx := i
		for chars[markerCloseIdx] != ')' {
			markerCloseIdx++
		}
		repetitions, _ := strconv.Atoi(string(chars[xIdx+1 : markerCloseIdx]))

		charsToRepeat := chars[markerCloseIdx+1 : markerCloseIdx+1+numChars]

		if strings.ContainsRune(string(charsToRepeat), '(') {
			count += decompress(charsToRepeat) * repetitions
		} else {
			count += numChars * repetitions
		}
		chars = chars[markerCloseIdx+1+numChars:]
	}
	return count
}
