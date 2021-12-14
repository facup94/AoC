package day11

import (
	"fmt"
	"regexp"
	"strings"

	"AoC/commons"
)

type solutionStep struct {
	building building
	step     int
}

type building struct {
	elevator      int
	components    uint64
	numComponents int
}

type componentGroup struct {
	name       rune
	components map[string]component
}

type component struct {
	name    string
	subtype string
	floor   int
}

func getComponentsFromInput(input []string) (components map[rune]componentGroup) {
	components = make(map[rune]componentGroup)

	reG := regexp.MustCompile(`\w+ generator`)
	reM := regexp.MustCompile(`\w+-\w+ microchip`)

	generatorMap := make(map[rune]int)
	microchipMap := make(map[rune]int)
	for floorNum, s := range input {
		if strings.Split(s, " ")[4] == "nothing" {
			continue
		}

		for _, ss := range reG.FindAllString(s, -1) {
			generatorMap[[]rune(ss)[0]] = floorNum
		}
		for _, ss := range reM.FindAllString(s, -1) {
			microchipMap[[]rune(ss)[0]] = floorNum
		}
	}

	for k, genFloor := range generatorMap {
		chipFloor := microchipMap[k]
		components[k-32] = componentGroup{
			name: k - 32,
			components: map[string]component{
				"microchip": {name: fmt.Sprintf("%cM", k-32), subtype: "microchip", floor: chipFloor},
				"generator": {name: fmt.Sprintf("%cM", k-32), subtype: "generator", floor: genFloor},
			},
		}
	}
	return components
}

func solve(initialComponents map[rune]componentGroup) (interface{}, error) {
	alreadyEvaluatedBuildings := make(map[uint64]bool)

	possiblesSteps := []solutionStep{{building: createInitialBuilding(initialComponents)}}
	step := possiblesSteps[0]

	for !step.building.hasAllComponentsFloor4() {
		buildingKey := step.building.toMapKey()

		if _, found := alreadyEvaluatedBuildings[buildingKey]; found {
			step = possiblesSteps[0]
			possiblesSteps = possiblesSteps[1:]
			continue
		} else {
			alreadyEvaluatedBuildings[buildingKey] = true
		}

		if !step.building.isValid() {
			step = possiblesSteps[0]
			possiblesSteps = possiblesSteps[1:]
			continue
		}

		possiblesSteps = append(possiblesSteps, step.generatePossibleSteps()...)

		step = possiblesSteps[0]
		possiblesSteps = possiblesSteps[1:]
	}

	return step.step, nil
}

func createInitialBuilding(components map[rune]componentGroup) (b building) {
	b.numComponents = len(components)

	var i int
	for _, group := range components {
		for _, comp := range group.components {
			switch comp.subtype {
			case "microchip":
				b.components |= uint64(comp.floor) << ((i + len(components)) * 2)
			case "generator":
				b.components |= uint64(comp.floor) << (i * 2)
			}
		}
		i++
	}

	return b
}

func (b building) hasAllComponentsFloor4() bool {
	if b.elevator != 3 {
		return false
	}

	for i := 0; i < b.numComponents*2; i++ {
		if b.components&(3<<(i*2))>>(i*2) != 3 {
			return false
		}
	}

	return true
}

func (b building) toMapKey() (key uint64) {
	key |= b.components
	key |= uint64(b.elevator) << (b.numComponents * 4)
	return key
}

func (b building) isValid() bool {
	for i := 0; i < b.numComponents; i++ {
		microchipFloor := b.components & (3 << ((i + b.numComponents) * 2)) >> ((i + b.numComponents) * 2)
		generatorFloor := b.components & (3 << (i * 2)) >> (i * 2)

		// Microchip and generator same floor = shield
		if microchipFloor == generatorFloor {
			continue
		}

		// If any other generator on same floor as microchip, then the chip fries
		for j := 0; j < b.numComponents; j++ {
			generatorFloor = b.components & (3 << (j * 2)) >> (j * 2)
			if microchipFloor == generatorFloor {
				return false
			}
		}
	}

	return true
}

func (s solutionStep) generatePossibleSteps() (steps []solutionStep) {
	for newElevatorFloor := commons.IntMax(0, s.building.elevator-1); newElevatorFloor <= commons.IntMin(3, s.building.elevator+1); newElevatorFloor++ {
		if newElevatorFloor == s.building.elevator {
			continue
		}

		var elevatorFloorRepeated uint64
		for i := 0; i < s.building.numComponents*2; i++ {
			elevatorFloorRepeated |= uint64(newElevatorFloor) << (i * 2)
		}

		itemsOnElevatorFloor := s.building.componentsOnElevatorFloor()

		for i, comp1 := range itemsOnElevatorFloor {
			for j, comp2 := range itemsOnElevatorFloor {
				if i < j {
					continue
				}

				newComponents := s.building.components

				newComponents &= ^comp1
				newComponents &= ^comp2

				newComponents |= comp1 & elevatorFloorRepeated
				newComponents |= comp2 & elevatorFloorRepeated

				b := building{
					elevator:      newElevatorFloor,
					components:    newComponents,
					numComponents: s.building.numComponents,
				}
				if !b.isValid() {
					continue
				}

				steps = append(steps, solutionStep{building: b, step: s.step + 1})
			}
		}
	}

	return steps
}

func (b building) componentsOnElevatorFloor() (components []uint64) {
	for i := 0; i < b.numComponents*2; i++ {
		pos := uint64(3 << (i * 2))
		if b.components&pos>>(i*2) == uint64(b.elevator) {
			components = append(components, pos)
		}
	}
	return components
}
