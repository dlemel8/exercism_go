package grains

import (
	"errors"
	"math"
)

const testVersion = 1

func Square(n int) (uint64, error) {
	if 1 > n || n > 64 {
		return 0, errors.New("invalid")
	}
	return 1 << uint(n-1), nil
}

func Total() (res uint64) {
	return math.MaxUint64
}
