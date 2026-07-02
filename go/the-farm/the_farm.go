package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	numCows int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numCows, e.message)
}

func DivideFood(calculator FodderCalculator, numCows int) (float64, error) {
	fodder, err := calculator.FodderAmount(numCows)
	if err != nil {
		return 0.0, err
	}

	fatteningFactor, err := calculator.FatteningFactor()
	if err != nil {
		return 0.0, err
	}

	return fodder * fatteningFactor / float64(numCows), nil
}

func ValidateInputAndDivideFood(calculator FodderCalculator, numCows int) (float64, error) {
	if numCows > 0 {
		return DivideFood(calculator, numCows)
	}

	return 0.0, errors.New("invalid number of cows")
}

func ValidateNumberOfCows(numCows int) error {
	switch {
	case numCows < 0:
		return &InvalidCowsError{
			numCows: numCows,
			message: "there are no negative cows",
		}
	case numCows == 0:
		return &InvalidCowsError{
			numCows: numCows,
			message: "no cows don't need food",
		}
	}

	return nil
}
