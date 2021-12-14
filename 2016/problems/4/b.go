package day4

import (
	"errors"
)

func B(input []string) (interface{}, error) {
	for _, s := range input {
		r := roomFromInput(s)
		if "northpole object storage" == r.decryptName() {
			return r.sectorID, nil
		}
	}

	return -1, errors.New("storage not found")
}

func (r room) decryptName() string {
	runes := []rune(r.encryptedName)
	result := make([]rune, len(runes))
	for i, nr := range runes {
		if nr == '-' {
			result[i] = ' '
		} else {
			result[i] = caesar(nr, r.sectorID)
		}
	}

	return string(result)
}

// caesar shifts characters by specified number of places.
// If beyond range, shift backward or forward.
func caesar(r rune, shift int) rune {
	s := int(r) + shift

	for s > 'z' {
		s -= 26
	}
	for s < 'a' {
		s += 26
	}

	return rune(s)
}
