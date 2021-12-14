package commons

func IntAbs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func IntMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func IntMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func RuneTernaryOperator(f bool, a, b rune) rune {
	if f {
		return a
	}
	return b
}

func Uint64Min(a, b uint64) uint64 {
	if a <= b {
		return a
	}
	return b
}

func Uint64Max(a, b uint64) uint64 {
	if a >= b {
		return a
	}
	return b
}
