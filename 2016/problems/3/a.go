package day3

import (
	"strconv"
	"strings"
)

func A(input []string) (interface{}, error) {
	var cant int
	for _, s := range input {
		ss := strings.Fields(s)
		a, _ := strconv.Atoi(ss[0])
		b, _ := strconv.Atoi(ss[1])
		c, _ := strconv.Atoi(ss[2])

		t := triangle{sideA: a, sideB: b, sideC: c}
		if t.isPossible() {
			cant++
		}
	}

	return cant, nil
}
