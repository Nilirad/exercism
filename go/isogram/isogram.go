package isogram

func IsIsogram(word string) bool {
	for i := range len(word) {
		charA := word[i]
		charA, isAlpha := validateAndNormalize(charA)
		if !isAlpha {
			continue
		}

		for j := i + 1; j < len(word); j++ {
			charB := word[j]
			charB, isAlpha = validateAndNormalize(charB)
			if !isAlpha {
				continue
			}

			if charA == charB {
				return false
			}
		}
	}

	return true
}

// validateAndNormalize ensures the byte is an ASCII letter
// and converts it to lowercase.
func validateAndNormalize(c byte) (byte, bool) {
	if c >= 'A' && c <= 'Z' {
		c += 'a' - 'A'
	}
	if c < 'a' || c > 'z' {
		return 0, false
	}

	return c, true
}
