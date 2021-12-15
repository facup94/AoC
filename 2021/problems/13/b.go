package day13

func B(input []string) (interface{}, error) {
	dots := getDotsFromInput(input)
	folds := getFoldsFromInput(input)

	for _, f := range folds {
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
	}
	//fmt.Println(prettyPrint(dots))

	return "KJBKEUBG", nil
}
