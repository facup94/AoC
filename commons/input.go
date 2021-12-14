package commons

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput(year, day int, test bool) ([]string, error) {
	path := fmt.Sprintf("./%d/problems/%d/input.txt", year, day)
	if test {
		path = fmt.Sprintf("./%d/problems/%d/input_test.txt", year, day)
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var inputStringSlice []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputStringSlice = append(inputStringSlice, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return inputStringSlice, nil
}
