package day12

import (
	"strings"
)

func A(input []string) (interface{}, error) {
	connections := make(map[string][]string)
	for _, s := range input {
		ss := strings.Split(s, "-")
		connections[ss[0]] = append(connections[ss[0]], ss[1])
		connections[ss[1]] = append(connections[ss[1]], ss[0])
	}

	var count int

	var paths = [][]string{{"start"}}
	for {
		if len(paths) == 0 {
			break
		}
		p := paths[0]
		paths = paths[1:]

		if p[len(p)-1] == "end" {
			count++
			continue
		}

		for _, nextCave := range connections[p[len(p)-1]] {
			if nextCave == "start" {
				continue
			}

			newPath := make([]string, len(p)+1)
			copy(newPath, p)
			newPath[len(newPath)-1] = nextCave
			if isSmallCave(nextCave) && !smallCavesOnlyVisitedOnce(newPath) {
				continue
			}

			paths = append(paths, newPath)
		}
	}

	return count, nil
}

func smallCavesOnlyVisitedOnce(path []string) bool {
	smallCaves := make(map[string]bool)
	for _, s := range path {
		if isSmallCave(s) {
			if _, present := smallCaves[s]; present {
				return false
			}
			smallCaves[s] = true
		}
	}
	return true
}
