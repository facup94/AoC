package year2019

import (
	day01 "AoC/2019/problems/1"
	day10 "AoC/2019/problems/10"
	day11 "AoC/2019/problems/11"
	day12 "AoC/2019/problems/12"
	day13 "AoC/2019/problems/13"
	day14 "AoC/2019/problems/14"
	day15 "AoC/2019/problems/15"
	day16 "AoC/2019/problems/16"
	day17 "AoC/2019/problems/17"
	day18 "AoC/2019/problems/18"
	day19 "AoC/2019/problems/19"
	day02 "AoC/2019/problems/2"
	day20 "AoC/2019/problems/20"
	day21 "AoC/2019/problems/21"
	day22 "AoC/2019/problems/22"
	day23 "AoC/2019/problems/23"
	day24 "AoC/2019/problems/24"
	day25 "AoC/2019/problems/25"
	day03 "AoC/2019/problems/3"
	day04 "AoC/2019/problems/4"
	day05 "AoC/2019/problems/5"
	day06 "AoC/2019/problems/6"
	day07 "AoC/2019/problems/7"
	day08 "AoC/2019/problems/8"
	day09 "AoC/2019/problems/9"
)

func GetPart(day int, part string) func([]string) (int, error) {
	switch day {
	case 1:
		return choosePart(day01.A, day01.B, part)
	case 2:
		return choosePart(day02.A, day02.B, part)
	case 3:
		return choosePart(day03.A, day03.B, part)
	case 4:
		return choosePart(day04.A, day04.B, part)
	case 5:
		return choosePart(day05.A, day05.B, part)
	case 6:
		return choosePart(day06.A, day06.B, part)
	case 7:
		return choosePart(day07.A, day07.B, part)
	case 8:
		return choosePart(day08.A, day08.B, part)
	case 9:
		return choosePart(day09.A, day09.B, part)
	case 10:
		return choosePart(day10.A, day10.B, part)
	case 11:
		return choosePart(day11.A, day11.B, part)
	case 12:
		return choosePart(day12.A, day12.B, part)
	case 13:
		return choosePart(day13.A, day13.B, part)
	case 14:
		return choosePart(day14.A, day14.B, part)
	case 15:
		return choosePart(day15.A, day15.B, part)
	case 16:
		return choosePart(day16.A, day16.B, part)
	case 17:
		return choosePart(day17.A, day17.B, part)
	case 18:
		return choosePart(day18.A, day18.B, part)
	case 19:
		return choosePart(day19.A, day19.B, part)
	case 20:
		return choosePart(day20.A, day20.B, part)
	case 21:
		return choosePart(day21.A, day21.B, part)
	case 22:
		return choosePart(day22.A, day22.B, part)
	case 23:
		return choosePart(day23.A, day23.B, part)
	case 24:
		return choosePart(day24.A, day24.B, part)
	case 25:
		return choosePart(day25.A, day25.B, part)
	}

	return nil
}

func choosePart(a, b func([]string) (int, error), part string) func([]string) (int, error) {
	if part == "A" {
		return a
	}
	return b
}
