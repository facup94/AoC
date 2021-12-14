package day5

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func B(input []string) (interface{}, error) {
	var x int
	password := make([]uint8, 8)
	positionsSet := [8]bool{false, false, false, false, false, false, false, false}

	for !positionsSet[0] || !positionsSet[1] || !positionsSet[2] || !positionsSet[3] || !positionsSet[4] || !positionsSet[5] || !positionsSet[6] || !positionsSet[7] {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(input[0]+strconv.Itoa(x))))

		if hash[0] == '0' && hash[1] == '0' && hash[2] == '0' && hash[3] == '0' && hash[4] == '0' {
			if hash[5] >= '0' && hash[5] <= '7' {
				if !positionsSet[hash[5]-'0'] {
					password[hash[5]-'0'] = hash[6]
					positionsSet[hash[5]-'0'] = true
				}
			}
		}

		x++
	}

	return string(password), nil
}
