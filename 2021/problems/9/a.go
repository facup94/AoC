package day09

func A(input []string) (interface{}, error) {
	heightmap := make([][]int, len(input))

	for y, s := range input {
		heightmap[y] = make([]int, len(s))
		for x, i := range s {
			heightmap[y][x] = int(i - 48)
		}
	}

	var sum int
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
				sum += 1 + heightmap[y][x]
			}
		}
	}

	return sum, nil
}
