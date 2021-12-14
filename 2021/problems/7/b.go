package day07

import (
	"AoC/commons"
)

func B(input []string) (interface{}, error) {
	crabs := getCrabsFromString(input[0])
	maxPos := commons.IntSliceMax(crabs)
	minPos := commons.IntSliceMin(crabs)

	var minFuelUsage = -1
	for pos := minPos; pos <= maxPos; pos++ {
		fuelCosts := make([]int, len(crabs))
		for i, crabPos := range crabs {
			distance := commons.IntAbs(crabPos - pos)
			for x := 1; x <= distance; x++ {
				fuelCosts[i] += x
			}
		}

		totalCost := commons.IntSliceSum(fuelCosts)
		if minFuelUsage == -1 || totalCost < minFuelUsage {
			minFuelUsage = totalCost
		}
	}

	return minFuelUsage, nil
}
