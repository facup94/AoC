package day8

import (
	"strconv"
	"strings"
)

type screen [6][50]bool

type instruction func(screen) screen

func stringToInstruction(s string) instruction {
	ss1 := strings.SplitN(s, " ", 2)
	switch ss1[0] {
	case "rect":
		ss2 := strings.Split(ss1[1], "x")
		w, _ := strconv.Atoi(ss2[0])
		h, _ := strconv.Atoi(ss2[1])

		return func(s screen) screen {
			for y := 0; y < h; y++ {
				for x := 0; x < w; x++ {
					s[y][x] = true
				}
			}
			return s
		}
	case "rotate":
		ss2 := strings.Split(ss1[1], " ")
		rowcolN, _ := strconv.Atoi(strings.Split(ss2[1], "=")[1])
		by, _ := strconv.Atoi(ss2[3])
		switch ss2[0] {
		case "row":
			by %= 50

			return func(s screen) screen {
				rowCopy := s[rowcolN]
				for i := 0; i < 50; i++ {
					s[rowcolN][(i+by)%50] = rowCopy[i]
				}
				return s
			}
		case "column":
			by %= 6

			return func(s screen) screen {
				var colCopy [6]bool
				for i := 0; i < 6; i++ {
					colCopy[i] = s[i][rowcolN]
				}

				for i := 0; i < 6; i++ {
					s[(i+by)%6][rowcolN] = colCopy[i]
				}
				return s
			}
		}

	}

	return func(s screen) screen {
		return screen{}
	}
}

func (s screen) String() string {
	var sb strings.Builder
	for _, row := range s {
		for _, pixel := range row {
			if pixel {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}
