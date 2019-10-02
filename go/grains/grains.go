package grains

import (
	"errors"
)

// Square returns the number of rice grains on a square
func Square(i int) (uint64, error) {
	if i > 64 || i < 1 {
		return uint64(0), errors.New("Cannot request a square number higher than 64 or less than 1")
	}
	return 1 << (i - 1), nil
}

// Total returns the total number of rice grains on the board
func Total() uint64 {
	return (1 << 64) - 1
}
