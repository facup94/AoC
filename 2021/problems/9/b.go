package day09

import (
	"fmt"
	"sort"
)

func B(input []string) (interface{}, error) {
	heightmap := make([][]int, len(input))

	for y, s := range input {
		heightmap[y] = make([]int, len(s))
		for x, i := range s {
			heightmap[y][x] = int(i - 48)
		}
	}

	var localMinors []position
	for y := 0; y < len(heightmap); y++ {
		for x := 0; x < len(heightmap[y]); x++ {
			localMinor := true

			if y-1 >= 0 {
				if heightmap[y][x] >= heightmap[y-1][x] {
					localMinor = false
				}
			}
			if x-1 >= 0 {
				if heightmap[y][x] >= heightmap[y][x-1] {
					localMinor = false
				}
			}
			if x+1 < len(heightmap[y]) {
				if heightmap[y][x] >= heightmap[y][x+1] {
					localMinor = false
				}
			}
			if y+1 < len(heightmap) {
				if heightmap[y][x] >= heightmap[y+1][x] {
					localMinor = false
				}
			}

			if localMinor {
				localMinors = append(localMinors, position{y, x})
			}
		}
	}

	var sinkSizes []int

	for _, sink := range localMinors {
		var count int
		alreadyChecked := make(map[string]bool)
		inBasin := []position{sink}

		for {
			if len(inBasin) == 0 {
				break
			}

			y, x := inBasin[0].y, inBasin[0].x
			locationKey := fmt.Sprintf("%v|%v", y, x)
			if checked := alreadyChecked[locationKey]; checked || heightmap[y][x] == 9 {
				inBasin = inBasin[1:]
				continue
			} else {
				alreadyChecked[locationKey] = true
			}

			count++

			if y-1 >= 0 && heightmap[y-1][x] > heightmap[y][x] {
				inBasin = append(inBasin, position{y - 1, x})
			}
			if x+1 < len(heightmap[0]) && heightmap[y][x+1] > heightmap[y][x] {
				inBasin = append(inBasin, position{y, x + 1})
			}
			if y+1 < len(heightmap) && heightmap[y+1][x] > heightmap[y][x] {
				inBasin = append(inBasin, position{y + 1, x})
			}
			if x-1 >= 0 && heightmap[y][x-1] > heightmap[y][x] {
				inBasin = append(inBasin, position{y, x - 1})
			}

			inBasin = inBasin[1:]
		}
		sinkSizes = append(sinkSizes, count)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sinkSizes)))

	return sinkSizes[0] * sinkSizes[1] * sinkSizes[2], nil
}

type position struct {
	y int
	x int
}
