package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fd FodderCalculator, cows int) (float64, error) {
	fa, err := fd.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	fm, err := fd.FatteningFactor()
	if err != nil {
		return 0, err
	}

	fi := float64(cows)

	return (fa * fm) / fi, nil

}

func ValidateInputAndDivideFood(fd FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(fd, cows)
	}

	return 0, errors.New("invalid number of cows")

}

type InvalidCowsError struct {
	numberOfCows int
	message      string
}

func (ic *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", ic.numberOfCows, ic.message)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			cows,
			"there are no negative cows",
		}
	} else if cows == 0 {
		return &InvalidCowsError{
			cows,
			"no cows don't need food",
		}
	} else {
		return nil
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
