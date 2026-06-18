package luhn

import "strings"

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}

	sum := 0

	mustDouble := len(id)%2 == 0
	for i := range len(id) {
		if id[i] < '0' || id[i] > '9' {
			return false
		}

		n := int(id[i] - '0')
		if mustDouble {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}

		sum += n
		mustDouble = !mustDouble
	}

	return sum%10 == 0
}
