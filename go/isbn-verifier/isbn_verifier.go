package isbnverifier

func IsValidISBN(isbn string) bool {
	multiplier := uint(10)
	accumulator := uint(0)
	for i := range len(isbn) {
		char := isbn[i]
		switch {
		case char == '-':
			continue
		case char == 'X':
			if multiplier != 1 {
				return false
			}
			accumulator += 10
			multiplier--
		case char >= '0' && char <= '9':
			digit := uint(char - '0')
			accumulator += digit * multiplier
			multiplier--
		default:
			return false
		}
	}

	numDigitsOk := multiplier == 0
	return numDigitsOk && accumulator%11 == 0
}
