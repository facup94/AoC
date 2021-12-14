package day04

import (
	"errors"
	"strconv"
	"strings"
)

func A(input []string) (interface{}, error) {
	boards, err := inputToBoards(append(input[2:], ""))
	if err != nil {
		return -1, err
	}

	drawnNumbersStr := strings.Split(input[0], ",")
	drawnNumbers := make([]int, len(drawnNumbersStr))
	for i, s := range drawnNumbersStr {
		n, err := strconv.Atoi(s)
		if err != nil {
			return -1, err
		}

		drawnNumbers[i] = n
	}

	for _, n := range drawnNumbers {
		for i := 0; i < len(boards); i++ {
			boards[i].setDrawnNumberTrue(n)
			if boards[i].hasBingo() {
				return boards[i].sumOfUnmarkedNumbers() * n, nil
			}
		}
	}

	return -1, errors.New("nobody won")
}
