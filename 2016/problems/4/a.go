package day4

import (
	"fmt"
	"sort"
	"strings"
)

func A(input []string) (interface{}, error) {
	var sum int

	for _, s := range input {
		r := roomFromInput(s)
		if r.isReal() {
			sum += r.sectorID
		}
	}

	return sum, nil
}

func (r room) isReal() bool {
	sortedLetters := sortNameLettersByCount(strings.ReplaceAll(r.encryptedName, "-", ""))
	for i := 0; i < 5; i++ {
		if sortedLetters[i] != r.checksum[i] {
			return false
		}
	}
	return true
}

type letterRepeatCounter struct {
	letter rune
	count  int
}

func sortNameLettersByCount(s string) []rune {
	countersMap := make(map[rune]letterRepeatCounter)
	for _, r := range []rune(s) {
		if _, exists := countersMap[r]; !exists {
			countersMap[r] = letterRepeatCounter{letter: r}
		}
		countersMap[r] = letterRepeatCounter{letter: r, count: countersMap[r].count + 1}
	}

	countersArr := make([]letterRepeatCounter, 0, len(countersMap))
	for _, c := range countersMap {
		countersArr = append(countersArr, c)
	}

	sort.Slice(countersArr, func(i, j int) bool {
		if countersArr[i].count < countersArr[j].count {
			return false
		}
		if countersArr[i].count == countersArr[j].count {
			return countersArr[i].letter < countersArr[j].letter
		}
		return true
	})

	sortedLetters := make([]rune, len(countersArr))
	for i, c := range countersArr {
		sortedLetters[i] = c.letter
	}

	return sortedLetters
}

func (l letterRepeatCounter) String() string {
	return fmt.Sprintf("%c (%d)", l.letter, l.count)
}
