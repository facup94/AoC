package day13

import (
	"fmt"
)

func A(input []string) (interface{}, error) {
	dots := getDotsFromInput(input)
	folds := getFoldsFromInput(input)

	//fmt.Println(prettyPrint(dots))

	f := folds[0]
	switch f.direction {
	case "y":
		for i := 0; i < len(dots); i++ {
			d := dots[i]
			if d.y < f.position {
				continue
			}

			distToFold := d.y - f.position
			d.y = f.position - distToFold

			dots[i] = d
		}
	case "x":
		for i := 0; i < len(dots); i++ {
			d := dots[i]
			if d.x < f.position {
				continue
			}

			distToFold := d.x - f.position
			d.x = f.position - distToFold

			dots[i] = d
		}
	}

	//fmt.Println(prettyPrint(dots))

	dotsMap := make(map[string]bool)
	for _, d := range dots {
		dotsMap[fmt.Sprintf("%d|%d", d.x, d.y)] = true
	}

	return len(dotsMap), nil
}
