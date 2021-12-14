package day07

import (
	"strconv"
	"strings"
)

func getCrabsFromString(s string) (crabs []int) {
	for _, ss := range strings.Split(s, ",") {
		v, _ := strconv.Atoi(ss)
		crabs = append(crabs, v)
	}
	return crabs
}
