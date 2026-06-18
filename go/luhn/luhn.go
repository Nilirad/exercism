package luhn

func Valid(id string) bool {
	var sum, count uint

	for i := len(id) - 1; i >= 0; i-- {
		char := id[i]
		switch {
		case char == ' ':
			continue
		case char >= '0' && char <= '9':
			digit := uint(char - '0')
			if count&1 == 1 {
				digit *= 2
				if digit > 9 {
					digit -= 9
				}
			}
			sum += digit
			count++
		default:
			return false
		}
	}

	return count > 1 && sum%10 == 0
}
