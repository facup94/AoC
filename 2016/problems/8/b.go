package day8

func B(input []string) (interface{}, error) {
	var scrn screen
	for _, s := range input {
		instr := stringToInstruction(s)
		scrn = instr(scrn)
	}

	return scrn.GetLetters(), nil
}

func (s screen) GetLetters() string {
	var letters []rune
	for letterN := 0; letterN < 50/5; letterN++ {
		var letter [6][5]bool
		for y := 0; y < 6; y++ {
			for x := 0; x < 5; x++ {
				letter[y][x] = s[y][x+(letterN*5)]
			}
		}

		for l, pixels := range letterDisplayFormat {
			if pixels == letter {
				letters = append(letters, l)
				break
			}
		}
	}

	return string(letters)
}

// I did this after printing screen
var letterDisplayFormat = map[rune][6][5]bool{
	'A': {{}, {}, {}, {}, {}, {}},
	'B': {{}, {}, {}, {}, {}, {}},
	'C': {{}, {}, {}, {}, {}, {}},
	'D': {{}, {}, {}, {}, {}, {}},
	'E': {{true, true, true, true, false}, {true, false, false, false, false}, {true, true, true, false, false}, {true, false, false, false, false}, {true, false, false, false, false}, {true, true, true, true, false}},
	'F': {{true, true, true, true, false}, {true, false, false, false, false}, {true, true, true, false, false}, {true, false, false, false, false}, {true, false, false, false, false}, {true, false, false, false, false}},
	'G': {{}, {}, {}, {}, {}, {}},
	'H': {{}, {}, {}, {}, {}, {}},
	'I': {{false, true, true, true, false}, {false, false, true, false, false}, {false, false, true, false, false}, {false, false, true, false, false}, {false, false, true, false, false}, {false, true, true, true, false}},
	'J': {{false, false, true, true, false}, {false, false, false, true, false}, {false, false, false, true, false}, {false, false, false, true, false}, {true, false, false, true, false}, {false, true, true, false, false}},
	'K': {{true, false, false, true, false}, {true, false, true, false, false}, {true, true, false, false, false}, {true, false, true, false, false}, {true, false, true, false, false}, {true, false, false, true, false}},
	'L': {{}, {}, {}, {}, {}, {}},
	'M': {{}, {}, {}, {}, {}, {}},
	'N': {{}, {}, {}, {}, {}, {}},
	'O': {{}, {}, {}, {}, {}, {}},
	'P': {{}, {}, {}, {}, {}, {}},
	'Q': {{}, {}, {}, {}, {}, {}},
	'R': {{true, true, true, false, false}, {true, false, false, true, false}, {true, false, false, true, false}, {true, true, true, false, false}, {true, false, true, false, false}, {true, false, false, true, false}},
	'S': {{}, {}, {}, {}, {}, {}},
	'T': {{}, {}, {}, {}, {}, {}},
	'U': {{}, {}, {}, {}, {}, {}},
	'V': {{}, {}, {}, {}, {}, {}},
	'W': {{}, {}, {}, {}, {}, {}},
	'X': {{}, {}, {}, {}, {}, {}},
	'Y': {{true, false, false, false, true}, {true, false, false, false, true}, {false, true, false, true, false}, {false, false, true, false, false}, {false, false, true, false, false}, {false, false, true, false, false}},
	'Z': {{}, {}, {}, {}, {}, {}},
}
