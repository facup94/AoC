package day7

func A(input []string) (interface{}, error) {
	var count int
	for _, s := range input {
		i := stringToIP(s)
		if i.hasABBAOnSupernet() && !i.hasABBAOnHypernet() {
			count++
		}
	}

	return count, nil
}
