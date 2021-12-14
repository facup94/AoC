package day01

import "AoC/commons"

func B(input []string) (interface{}, error) {
	measurements, err := commons.StringSliceToIntSlice(input)
	if err != nil {
		return -1, err
	}

	var increments int
	for i := 0; i < len(measurements)-3; i++ {
		if measurements[i] < measurements[i+3] {
			increments++
		}
	}

	return increments, nil
}
