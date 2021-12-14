package day07

import (
	"strconv"
	"strings"

	"AoC/commons"
)

func A(input []string) (interface{}, error) {
	var crabs []int
	for _, s := range strings.Split(input[0], ",") {
		v, _ := strconv.Atoi(s)
		crabs = append(crabs, v)
	}

	maxPos := commons.IntSliceMax(crabs)
	minPos := commons.IntSliceMin(crabs)

	var minFuelUsage = -1
	for pos := minPos; pos <= maxPos; pos++ {
		fuelCosts := make([]int, len(crabs))
		for i, crabPos := range crabs {
			fuelCosts[i] = commons.IntAbs(crabPos - pos)
		}

		totalCost := commons.IntSliceSum(fuelCosts)
		if minFuelUsage == -1 || totalCost < minFuelUsage {
			minFuelUsage = totalCost
		}
	}

	return minFuelUsage, nil
}
