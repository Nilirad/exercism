// Tools for calculating the hamming distance.
package hamming

import (
	"errors"
	"fmt"
)

// Caluculates the hamming distance between two DNA sequences.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("Sequences are of unequal lenght. len(a)=%d, len(b)=%d", len(a), len(b))
	}

	length := len(a)
	distance := 0

	for i := range length {
		if !isLetterDna(a[i]) {
			return 0, errors.New("String `a` is not a DNA sequence.")
		}
		if !isLetterDna(b[i]) {
			return 0, errors.New("String `b` is not a DNA sequence.")
		}

		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}

// Helper function to determine if a byte
// correspond to a DNA base.
func isLetterDna(letter byte) bool {
	switch letter {
	case 'A', 'T', 'G', 'C':
		return true
	default:
		return false
	}
}
