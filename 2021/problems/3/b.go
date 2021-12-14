package day03

import (
	"AoC/commons"
)

func B(input []string) (interface{}, error) {
	diagnosticReport, err := commons.StringBinarySliceToInt64Slice(input)
	if err != nil {
		return -1, err
	}

	O2RatingMeasurements := make([]int64, len(diagnosticReport))
	CO2RatingMeasurements := make([]int64, len(diagnosticReport))
	copy(O2RatingMeasurements, diagnosticReport)
	copy(CO2RatingMeasurements, diagnosticReport)

	pos := 0
	for len(O2RatingMeasurements) > 1 {
		zeroInPosition, oneInPosition := splitByZeroAndOneInPosition(O2RatingMeasurements, len(input[0])-1-pos)

		if len(oneInPosition) >= len(zeroInPosition) {
			O2RatingMeasurements = oneInPosition
		} else {
			O2RatingMeasurements = zeroInPosition
		}

		pos++
	}

	pos = 0
	for len(CO2RatingMeasurements) > 1 {
		zeroInPosition, oneInPosition := splitByZeroAndOneInPosition(CO2RatingMeasurements, len(input[0])-1-pos)

		if len(zeroInPosition) <= len(oneInPosition) {
			CO2RatingMeasurements = zeroInPosition
		} else {
			CO2RatingMeasurements = oneInPosition
		}

		pos++
	}

	return int(O2RatingMeasurements[0] * CO2RatingMeasurements[0]), nil
}

func splitByZeroAndOneInPosition(in []int64, pos int) (zeros, ones []int64) {
	for _, n := range in {
		if (n & (1 << pos)) > 0 {
			ones = append(ones, n)
		} else {
			zeros = append(zeros, n)
		}
	}
	return zeros, ones
}
