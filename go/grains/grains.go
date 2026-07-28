package grains

import (
	"fmt"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, fmt.Errorf("%d is not a chessboard square", number)
	}

	// Each bit in an integer corresponds a power of two;
	// therefore, an integer with all `0`s and one `1`
	// corresponds to a specific square of Sissa's Chessboard.
	return uint64(1) << (number - 1), nil
}

func Total() uint64 {
	// We have 64 squares, and 64 bits in the number.
	// Each bit of an integer
	// corresponds to a specific square of Sissa's Chessboard.
	// Setting all bits to `1` gives the total grain count.
	return math.MaxUint64
}
