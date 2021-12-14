package day12

import "strings"

func isSmallCave(caveName string) bool {
	return caveName == strings.ToLower(caveName)
}
