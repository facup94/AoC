package day3

type triangle struct {
	sideA int
	sideB int
	sideC int
}

func (t triangle) isPossible() bool {
	return t.sideA+t.sideB > t.sideC && t.sideA+t.sideC > t.sideB && t.sideB+t.sideC > t.sideA
}
