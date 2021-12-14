package day08

import (
	"strings"
)

const (
	segments1 = 2
	segments4 = 4
	segments7 = 3
	segments8 = 7
)

func A(input []string) (interface{}, error) {
	var count int
	for _, s := range input {
		for _, ss := range strings.Fields(strings.Split(s, " | ")[1]) {
			l := len(ss)
			if l == segments1 || l == segments4 || l == segments7 || l == segments8 {
				count++
			}
		}
	}

	return count, nil
}
