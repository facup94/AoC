package day08

import (
	"math"
	sortstd "sort"
	"strings"
)

func B(input []string) (interface{}, error) {
	var sum int

	for _, s := range input {
		var patterns [][]rune
		for _, ss := range strings.Fields(strings.Split(s, " | ")[0]) {
			patterns = append(patterns, sort([]rune(ss)))
		}
		var outputPatterns []string
		for _, ss := range strings.Fields(strings.Split(s, " | ")[1]) {
			outputPatterns = append(outputPatterns, string(sort([]rune(ss))))
		}

		digitSegments := make(map[int][]rune)
		var len5patterns [][]rune
		var len6patterns [][]rune

		// 1, 4, 7, 8
		for _, pattern := range patterns {
			switch len(pattern) {
			case 2:
				digitSegments[1] = pattern
			case 3:
				digitSegments[7] = pattern
			case 4:
				digitSegments[4] = pattern
			case 5:
				len5patterns = append(len5patterns, pattern)
			case 6:
				len6patterns = append(len6patterns, pattern)
			case 7:
				digitSegments[8] = pattern
			}
		}

		// 2
		for i, pattern := range len5patterns {
			if len(sub(pattern, add(digitSegments[7], digitSegments[4]))) == 2 {
				digitSegments[2] = pattern
				len5patterns = append(len5patterns[:i], len5patterns[i+1:]...)
				break
			}
		}

		// 0
		for i, pattern := range len6patterns {
			if len(sub(sub(digitSegments[4], digitSegments[1]), pattern)) > 0 {
				digitSegments[0] = pattern
				len6patterns = append(len6patterns[:i], len6patterns[i+1:]...)
				break
			}
		}

		// 6, 9
		for _, pattern := range len6patterns {
			if len(sub(pattern, digitSegments[4])) == 2 {
				digitSegments[9] = pattern
			} else if len(sub(pattern, digitSegments[4])) == 3 {
				digitSegments[6] = pattern
			}
		}

		// 3, 5
		for _, pattern := range len5patterns {
			if len(sub(pattern, digitSegments[1])) == 3 {
				digitSegments[3] = pattern
			} else if len(sub(pattern, digitSegments[1])) == 4 {
				digitSegments[5] = pattern
			}
		}

		segmentsDigits := make(map[string]int)
		for digit, segments := range digitSegments {
			segmentsDigits[string(segments)] = digit
		}

		var code int
		for i, pattern := range outputPatterns {
			code += segmentsDigits[pattern] * int(math.Pow(10, float64(3-i)))
		}

		sum += code
	}
	return sum, nil
}

func add(a, b []rune) []rune {
	patternMap := make(map[rune]bool)
	for _, r := range a {
		patternMap[r] = true
	}
	for _, r := range b {
		patternMap[r] = true
	}

	var pattern []rune
	for r := range patternMap {
		pattern = append(pattern, r)
	}
	return pattern
}

func sub(a, b []rune) []rune {
	patternMap := make(map[rune]bool)
	for _, r := range a {
		patternMap[r] = true
	}
	for _, r := range b {
		delete(patternMap, r)
	}

	var pattern []rune
	for r := range patternMap {
		pattern = append(pattern, r)
	}
	return pattern
}

func sort(pattern []rune) []rune {
	sortstd.Slice(pattern, func(i, j int) bool {
		return pattern[i] < pattern[j]
	})
	return pattern
}
