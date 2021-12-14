package day11

import (
	"AoC/commons"
)

func A(input []string) (interface{}, error) {
	octopuses := make([][]int, len(input))
	for i, s := range input {
		v, err := commons.RuneSliceToIntSlice([]rune(s))
		if err != nil {
			return nil, err
		}
		octopuses[i] = v
	}

	var count int
	for i := 0; i < 100; i++ {
		// Increment all
		for y := 0; y < len(octopuses); y++ {
			for x := 0; x < len(octopuses[y]); x++ {
				octopuses[y][x]++
			}
		}

		var someoneFlashes = true
		for someoneFlashes {
			someoneFlashes = false
			for y := 0; y < len(octopuses); y++ {
				for x := 0; x < len(octopuses[y]); x++ {
					if octopuses[y][x] > 9 {
						octopuses[y][x] = -10

						// -1, -1
						if y-1 >= 0 && x-1 >= 0 {
							octopuses[y-1][x-1]++
							if octopuses[y-1][x-1] > 9 {
								someoneFlashes = true
							}
						}
						// -1, 0
						if y-1 >= 0 {
							octopuses[y-1][x]++
							if octopuses[y-1][x] > 9 {
								someoneFlashes = true
							}
						}
						// -1, +1
						if y-1 >= 0 && x+1 < len(octopuses[y-1]) {
							octopuses[y-1][x+1]++
							if octopuses[y-1][x+1] > 9 {
								someoneFlashes = true
							}
						}
						// 0, -1
						if x-1 >= 0 {
							octopuses[y][x-1]++
							if octopuses[y][x-1] > 9 {
								someoneFlashes = true
							}
						}
						// 0, +1
						if x+1 < len(octopuses[y]) {
							octopuses[y][x+1]++
							if octopuses[y][x+1] > 9 {
								someoneFlashes = true
							}
						}
						// +1, -1
						if y+1 < len(octopuses) && x-1 >= 0 {
							octopuses[y+1][x-1]++
							if octopuses[y+1][x-1] > 9 {
								someoneFlashes = true
							}
						}
						// +1, 0
						if y+1 < len(octopuses) {
							octopuses[y+1][x]++
							if octopuses[y+1][x] > 9 {
								someoneFlashes = true
							}
						}
						// +1, +1
						if y+1 < len(octopuses) && x+1 < len(octopuses[y+1]) {
							octopuses[y+1][x+1]++
							if octopuses[y+1][x+1] > 9 {
								someoneFlashes = true
							}
						}
					}

				}
			}
		}

		//
		for y := 0; y < len(octopuses); y++ {
			for x := 0; x < len(octopuses[y]); x++ {
				if octopuses[y][x] < 0 {
					octopuses[y][x] = 0
					count++
				}
			}
		}
	}

	return count, nil
}
