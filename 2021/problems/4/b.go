package day04

import (
	"errors"
	"strconv"
	"strings"

	"AoC/commons"
)

func B(input []string) (interface{}, error) {
	drawnNumbersStr := strings.Split(input[0], ",")
	drawnNumbers := make([]int, len(drawnNumbersStr))
	for i, s := range drawnNumbersStr {
		n, err := strconv.Atoi(s)
		if err != nil {
			return -1, err
		}

		drawnNumbers[i] = n
	}

	boards, err := inputToBoards(append(input[2:], ""))
	if err != nil {
		return -1, err
	}

	boardsThatWon := make([]bool, len(boards))

	for _, n := range drawnNumbers {
		for i := 0; i < len(boards); i++ {
			boards[i].setDrawnNumberTrue(n)
			if boards[i].hasBingo() {
				boardsThatWon[i] = true
				if commons.BoolSliceAllTrue(boardsThatWon) {
					return boards[i].sumOfUnmarkedNumbers() * n, nil
				}
			}
		}
	}

	return -1, errors.New("not all boards won")
}
