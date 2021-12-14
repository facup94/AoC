package day5

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func A(input []string) (interface{}, error) {
	var password []uint8

	var x int
	for len(password) < 8 {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(input[0]+strconv.Itoa(x))))

		if hash[0] == '0' && hash[1] == '0' && hash[2] == '0' && hash[3] == '0' && hash[4] == '0' {
			password = append(password, hash[5])
		}

		x++
	}

	return string(password), nil
}
