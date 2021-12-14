package day03

import "AoC/commons"

func A(input []string) (interface{}, error) {
	diagnosticReport, err := commons.StringBinarySliceToInt64Slice(input)
	if err != nil {
		return -1, err
	}

	countOnes := make([]int, len(input[0]))
	for _, i := range diagnosticReport {
		for j := len(countOnes) - 1; j >= 0; j-- {
			if i&(1<<j) != 0 {
				countOnes[j]++
			}
		}
	}

	var gamaRate, epsilonRate int64
	for i, n := range countOnes {
		if n > len(diagnosticReport)/2 {
			gamaRate |= 1 << i
		} else {
			epsilonRate |= 1 << i
		}
	}

	return int(gamaRate * epsilonRate), nil
}
