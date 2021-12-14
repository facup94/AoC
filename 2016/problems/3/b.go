package day3

import (
	"strconv"
	"strings"
)

func B(input []string) (interface{}, error) {
	var cant int

	for i := 0; i < len(input); i += 3 {
		in1 := strings.Fields(input[i])
		in2 := strings.Fields(input[i+1])
		in3 := strings.Fields(input[i+2])

		for n := 0; n < 3; n++ {
			a, _ := strconv.Atoi(in1[n])
			b, _ := strconv.Atoi(in2[n])
			c, _ := strconv.Atoi(in3[n])

			t := triangle{sideA: a, sideB: b, sideC: c}
			if t.isPossible() {
				cant++
			}
		}
	}

	return cant, nil
}
