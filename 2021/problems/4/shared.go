package day04

import (
	"fmt"
	"strconv"
	"strings"
)

type boardNumber struct {
	number int
	drawn  bool
}

type board [5][5]boardNumber

func inputToBoards(input []string) ([]board, error) {
	boards := make([]board, len(input)/6)
	for x := 0; x < len(boards); x++ {
		for j, row := range input[x*6 : (x+1)*6-1] {
			for i, s := range strings.Fields(row) {
				n, err := strconv.Atoi(s)
				if err != nil {
					return nil, err
				}

				boards[x][j][i].number = n
			}
		}
	}

	return boards, nil
}

func (b board) hasBingo() bool {
	for _, row := range b {
		if sliceHasBingo(row) {
			return true
		}
	}
	for j := 0; j < len(b[0]); j++ {
		var column [5]boardNumber
		for i, row := range b {
			column[i] = row[j]
		}
		if sliceHasBingo(column) {
			return true
		}
	}
	return false
}

func sliceHasBingo(row [5]boardNumber) bool {
	for _, n := range row {
		if !n.drawn {
			return false
		}
	}
	return true
}

func (b *board) setDrawnNumberTrue(n int) {
	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			if b[j][i].number == n {
				b[j][i].drawn = true
			}
		}
	}
}

func (b board) sumOfUnmarkedNumbers() (sum int) {
	for _, row := range b {
		for _, boardNum := range row {
			if !boardNum.drawn {
				sum += boardNum.number
			}
		}
	}
	return sum
}

func (b board) String() string {
	var sb strings.Builder
	for _, row := range b {
		for _, bNum := range row {
			s := fmt.Sprintf("%02d", bNum.number)
			if bNum.drawn {
				sb.WriteRune('*')
				sb.WriteString(s)
				sb.WriteRune('*')
				//for _, r := range s {
				//	switch r {
				//	case 0:
				//		sb.WriteRune('𝟎')
				//	case 1:
				//		sb.WriteRune('𝟏')
				//	case 2:
				//		sb.WriteRune('𝟐')
				//	case 3:
				//		sb.WriteRune('𝟑')
				//	case 4:
				//		sb.WriteRune('𝟒')
				//	case 5:
				//		sb.WriteRune('𝟓')
				//	case 6:
				//		sb.WriteRune('𝟔')
				//	case 7:
				//		sb.WriteRune('𝟕')
				//	case 8:
				//		sb.WriteRune('𝟖')
				//	case 9:
				//		sb.WriteRune('𝟗')
				//	}
				//}
			} else {
				sb.WriteString(s)
			}
			sb.WriteRune(' ')
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}
