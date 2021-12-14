package commons

func BoolSliceAllTrue(slice []bool) bool {
	for _, b := range slice {
		if !b {
			return false
		}
	}
	return true
}

func StringSlicesHaveAtLeastOneItemInCommon(a, b []string) bool {
	for _, x := range a {
		for _, y := range b {
			if x == y {
				return true
			}
		}
	}
	return false
}

func IntSliceMax(in []int) (max int) {
	for i, v := range in {
		if i == 0 || v > max {
			max = v
		}
	}
	return max
}

func IntSliceMin(in []int) (min int) {
	for i, v := range in {
		if i == 0 || v < min {
			min = v
		}
	}
	return min
}

func IntSliceSum(in []int) (sum int) {
	for _, v := range in {
		sum += v
	}
	return sum
}
