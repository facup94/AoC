package day2

func B(input []string) (interface{}, error) {
	pad := keypadB{currentPos: '5'}

	var resultArr []rune
	for _, instructions := range input {
		for _, l := range instructions {
			pad.move(l)
		}

		resultArr = append(resultArr, pad.currentPos)
	}

	return string(resultArr), nil
}

type keypadB struct {
	currentPos rune
}

func (k *keypadB) move(direction int32) {
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

func (k *keypadB) moveUp() {
	switch k.currentPos {
	case '1':
		k.currentPos = '1'
	case '2':
		k.currentPos = '2'
	case '3':
		k.currentPos = '1'
	case '4':
		k.currentPos = '4'
	case '5':
		k.currentPos = '5'
	case '6':
		k.currentPos = '2'
	case '7':
		k.currentPos = '3'
	case '8':
		k.currentPos = '4'
	case '9':
		k.currentPos = '9'
	case 'A':
		k.currentPos = '6'
	case 'B':
		k.currentPos = '7'
	case 'C':
		k.currentPos = '8'
	case 'D':
		k.currentPos = 'B'
	}
}

func (k *keypadB) moveDown() {
	switch k.currentPos {
	case '1':
		k.currentPos = '3'
	case '2':
		k.currentPos = '6'
	case '3':
		k.currentPos = '7'
	case '4':
		k.currentPos = '8'
	case '5':
		k.currentPos = '5'
	case '6':
		k.currentPos = 'A'
	case '7':
		k.currentPos = 'B'
	case '8':
		k.currentPos = 'C'
	case '9':
		k.currentPos = '9'
	case 'A':
		k.currentPos = 'A'
	case 'B':
		k.currentPos = 'D'
	case 'C':
		k.currentPos = 'C'
	case 'D':
		k.currentPos = 'D'
	}
}

func (k *keypadB) moveLeft() {
	switch k.currentPos {
	case '1':
		k.currentPos = '1'
	case '2':
		k.currentPos = '2'
	case '3':
		k.currentPos = '2'
	case '4':
		k.currentPos = '3'
	case '5':
		k.currentPos = '5'
	case '6':
		k.currentPos = '5'
	case '7':
		k.currentPos = '6'
	case '8':
		k.currentPos = '7'
	case '9':
		k.currentPos = '8'
	case 'A':
		k.currentPos = 'A'
	case 'B':
		k.currentPos = 'A'
	case 'C':
		k.currentPos = 'B'
	case 'D':
		k.currentPos = 'D'
	}
}

func (k *keypadB) moveRight() {
	switch k.currentPos {
	case '1':
		k.currentPos = '1'
	case '2':
		k.currentPos = '3'
	case '3':
		k.currentPos = '4'
	case '4':
		k.currentPos = '4'
	case '5':
		k.currentPos = '6'
	case '6':
		k.currentPos = '7'
	case '7':
		k.currentPos = '8'
	case '8':
		k.currentPos = '9'
	case '9':
		k.currentPos = '9'
	case 'A':
		k.currentPos = 'B'
	case 'B':
		k.currentPos = 'C'
	case 'C':
		k.currentPos = 'C'
	case 'D':
		k.currentPos = 'D'
	}
}
