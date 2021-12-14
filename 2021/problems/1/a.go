package day01

import "AoC/commons"

func A(input []string) (interface{}, error) {
	measures, err := commons.StringSliceToIntSlice(input)
	if err != nil {
		return -1, err
	}

	var increments int
	for i := 1; i < len(measures); i++ {
		if measures[i] > measures[i-1] {
			increments++
		}
	}

	return increments, nil
}
