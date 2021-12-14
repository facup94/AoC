package day10

func A(input []string) (interface{}, error) {
	var score int
	for _, s := range input {
		switch getFirstIllegalCharacter(s) {
		case ')':
			score += 3
		case ']':
			score += 57
		case '}':
			score += 1197
		case '>':
			score += 25137
		}
	}

	return score, nil
}

func getFirstIllegalCharacter(s string) rune {
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
				return r
			}
			openChunks = openChunks[:len(openChunks)-1]
		case ']':
			if openChunks[len(openChunks)-1] != ']' {
				return r
			}
			openChunks = openChunks[:len(openChunks)-1]
		case '}':
			if openChunks[len(openChunks)-1] != '}' {
				return r
			}
			openChunks = openChunks[:len(openChunks)-1]
		case '>':
			if openChunks[len(openChunks)-1] != '>' {
				return r
			}
			openChunks = openChunks[:len(openChunks)-1]
		}
	}

	if len(openChunks) > 0 {
		return 'I'
	}
	return 'C'
}
