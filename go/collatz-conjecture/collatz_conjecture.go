package collatzconjecture

import (
	"errors"
	"math"
)

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("`n` must be greater than zero.")
	}

	steps := 0

	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			if n >= (math.MaxInt-1)/3 {
				return 0, errors.New("Integer overflow")
			}
			n = n*3 + 1
		}
		steps++
	}

	return steps, nil
}
