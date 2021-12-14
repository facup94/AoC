package day10

import (
	"sort"
)

func B(input []string) (interface{}, error) {
	var scores []int
	for _, s := range input {
		missingChars := getMissingChars(s)
		if len(missingChars) > 0 {
			scores = append(scores, calculateScore(missingChars))
		}
	}

	sort.Ints(scores)

	return scores[len(scores)/2], nil
}

func getMissingChars(s string) []rune {
	var openChunks []rune
	for _, r := range []rune(s) {
		switch r {
		case '(':
			openChunks = append(openChunks, ')')
		case '[':
			openChunks = append(openChunks, ']')
		case '{':
			openChunks = append(openChunks, '}')
		case '<':
			openChunks = append(openChunks, '>')
		case ')':
			if openChunks[len(openChunks)-1] != ')' {
				return nil
			}
			openChunks = openChunks[:len(openChunks)-1]
		case ']':
			if openChunks[len(openChunks)-1] != ']' {
				return nil
			}
			openChunks = openChunks[:len(openChunks)-1]
		case '}':
			if openChunks[len(openChunks)-1] != '}' {
				return nil
			}
			openChunks = openChunks[:len(openChunks)-1]
		case '>':
			if openChunks[len(openChunks)-1] != '>' {
				return nil
			}
			openChunks = openChunks[:len(openChunks)-1]
		}
	}

	if len(openChunks) > 0 {
		for i := 0; i < len(openChunks)/2; i++ {
			openChunks[i], openChunks[len(openChunks)-1-i] = openChunks[len(openChunks)-1-i], openChunks[i]
		}
		return openChunks
	}
	return nil
}

func calculateScore(chars []rune) (score int) {
	for _, char := range chars {
		score *= 5
		switch char {
		case ')':
			score += 1
		case ']':
			score += 2
		case '}':
			score += 3
		case '>':
			score += 4
		}
	}
	return score
}
