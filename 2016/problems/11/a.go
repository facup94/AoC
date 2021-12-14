package day11

func A(input []string) (interface{}, error) {
	initialComponents := getComponentsFromInput(input)

	return solve(initialComponents)
}
