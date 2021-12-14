package day2

import (
	"math"
)

func A(input []string) (interface{}, error) {
	pad := keypadA{currentPos: 5}

	var resultArr []int
	for _, instructions := range input {
		for _, l := range instructions {
			pad.move(l)
		}

		resultArr = append(resultArr, pad.currentPos)
	}

	var result int
	for i, n := range resultArr {
		result += n * int(math.Pow(10, float64(len(resultArr)-1-i)))
	}
	return result, nil
}

type keypadA struct {
	currentPos int
}

func (k *keypadA) move(direction int32) {
	switch direction {
	case 'U':
		k.moveUp()
	case 'R':
		k.moveRight()
	case 'D':
		k.moveDown()
	case 'L':
		k.moveLeft()
	}
}

func (k *keypadA) moveUp() {
	switch k.currentPos {
	case 1:
		k.currentPos = 1
	case 2:
		k.currentPos = 2
	case 3:
		k.currentPos = 3
	case 4:
		k.currentPos = 1
	case 5:
		k.currentPos = 2
	case 6:
		k.currentPos = 3
	case 7:
		k.currentPos = 4
	case 8:
		k.currentPos = 5
	case 9:
		k.currentPos = 6
	}
}

func (k *keypadA) moveDown() {
	switch k.currentPos {
	case 1:
		k.currentPos = 4
	case 2:
		k.currentPos = 5
	case 3:
		k.currentPos = 6
	case 4:
		k.currentPos = 7
	case 5:
		k.currentPos = 8
	case 6:
		k.currentPos = 9
	case 7:
		k.currentPos = 7
	case 8:
		k.currentPos = 8
	case 9:
		k.currentPos = 9
	}
}

func (k *keypadA) moveLeft() {
	switch k.currentPos {
	case 1:
		k.currentPos = 1
	case 2:
		k.currentPos = 1
	case 3:
		k.currentPos = 2
	case 4:
		k.currentPos = 4
	case 5:
		k.currentPos = 4
	case 6:
		k.currentPos = 5
	case 7:
		k.currentPos = 7
	case 8:
		k.currentPos = 7
	case 9:
		k.currentPos = 8
	}
}

func (k *keypadA) moveRight() {
	switch k.currentPos {
	case 1:
		k.currentPos = 2
	case 2:
		k.currentPos = 3
	case 3:
		k.currentPos = 3
	case 4:
		k.currentPos = 5
	case 5:
		k.currentPos = 6
	case 6:
		k.currentPos = 6
	case 7:
		k.currentPos = 8
	case 8:
		k.currentPos = 9
	case 9:
		k.currentPos = 9
	}
}
