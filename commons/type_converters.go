package commons

import "strconv"

func StringSliceToIntSlice(src []string) ([]int, error) {
	dst := make([]int, len(src))
	for i, s := range src {
		if v, err := strconv.Atoi(s); err != nil {
			return nil, err
		} else {
			dst[i] = v
		}
	}
	return dst, nil
}

func StringBinarySliceToInt64Slice(src []string) ([]int64, error) {
	dst := make([]int64, len(src))
	for i, s := range src {
		v, err := strconv.ParseInt(s, 2, 64)
		if err != nil {
			return nil, err
		}
		dst[i] = v
	}
	return dst, nil
}

func RuneSliceToIntSlice(src []rune) ([]int, error) {
	dst := make([]int, len(src))
	for i, r := range src {
		dst[i] = int(r - 48)
	}
	return dst, nil
}
