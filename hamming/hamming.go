package hamming

import "errors"

const testVersion = 6

func Distance(a, b string) (int, error) {
	res := 0
	if len(a) != len(b) {
		return -1, errors.New("must be same length")
	}
	for i := range a {
		if a[i] != b[i] {
			res++
		}
	}
	return res, nil
}
