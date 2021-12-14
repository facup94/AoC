package main

import (
	"flag"
	"fmt"
	"time"

	year2016 "AoC/2016"
	year2021 "AoC/2021"
	"AoC/commons"
)

func main() {
	mode := flag.String("mode", "", "run mode")
	year := flag.Int("year", -1, "year to run")
	day := flag.Int("day", -1, "day to run")
	part := flag.String("part", "", "part of day to run")
	test := flag.Bool("test", false, "run on test input.txt")
	flag.Parse()

	switch *mode {
	case "today":
		runToday(*test)
	case "part":
		runPart(*year, *day, *part, *test)
	case "day":
		runDay(*year, *day, *test)
	case "year":
		runYear(*year, *test)
	default:
		runAll(*test)
	}
}

func runAll(test bool) {
	for i := 2019; i < 2022; i++ {
		runYear(i, test)
	}
}

func runYear(year int, test bool) {
	for i := 1; i < 26; i++ {
		runDay(year, i, test)
	}
}

func runToday(test bool) {
	t := time.Now()
	runDay(t.Year(), t.Day(), test)
}

func runDay(year, day int, test bool) {
	runPart(year, day, "A", test)
	runPart(year, day, "B", test)
}

func runPart(year, day int, part string, test bool) {
	input, _ := commons.ReadInput(year, day, test)

	var f func([]string) (interface{}, error)
	switch year {
	case 2016:
		f = year2016.GetPart(day, part)
	case 2021:
		f = year2021.GetPart(day, part)
	}

	executeSolution(year, day, part, func() (interface{}, error) { return f(input) })
}

func executeSolution(year, day int, part string, f func() (interface{}, error)) {
	msgFormat := "[Year:%d][Day:%d][Part:%s][Time:%s][Result:%s] Output: %v\n"

	start := time.Now()
	if result, err := f(); err == nil {
		fmt.Printf(msgFormat, year, day, part, time.Now().Sub(start), "Ok", result)
	} else {
		fmt.Printf(msgFormat, year, day, part, time.Now().Sub(start), "Error", err.Error())
	}
}
