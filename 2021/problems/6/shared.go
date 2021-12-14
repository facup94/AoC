package day06

import (
	"strconv"
	"strings"
)

func getFishesFromInput(s string) []int {
	var fishes []int
	for _, ss := range strings.Split(s, ",") {
		n, _ := strconv.Atoi(ss)
		fishes = append(fishes, n)
	}
	return fishes
}
