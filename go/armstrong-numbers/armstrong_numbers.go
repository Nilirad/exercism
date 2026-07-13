package armstrongnumbers

import "math"

func IsNumber(n int) bool {
	digits := int(math.Floor(math.Log10(float64(n)))) + 1

	sum := 0
	for i := range digits {
		digit := int(math.Floor(float64(n)/math.Pow10(i))) % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
	}

	return sum == n
}
