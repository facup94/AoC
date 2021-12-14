package day11

func B(input []string) (interface{}, error) {
	initialComponents := getComponentsFromInput(input)
	initialComponents['E'] = componentGroup{name: 'E', components: map[string]component{"EM": {name: "EM", subtype: "microchip"}, "EG": {name: "EG", subtype: "generator"}}}
	initialComponents['D'] = componentGroup{name: 'D', components: map[string]component{"DM": {name: "DM", subtype: "microchip"}, "DG": {name: "DG", subtype: "generator"}}}

	return solve(initialComponents)
}
