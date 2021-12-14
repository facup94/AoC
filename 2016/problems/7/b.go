package day7

import (
	"fmt"

	"AoC/commons"
)

func B(input []string) (interface{}, error) {
	var count int
	for _, s := range input {
		i := stringToIP(s)

		abas := i.getABAs()
		babs := i.getBABs()

		abasReversed := make([]string, len(abas))
		for x, aba := range abas {
			abasReversed[x] = fmt.Sprintf("%c%c%c", aba[1], aba[0], aba[1])
		}

		if commons.StringSlicesHaveAtLeastOneItemInCommon(abasReversed, babs) {
			count++
		}
	}

	return count, nil
}
