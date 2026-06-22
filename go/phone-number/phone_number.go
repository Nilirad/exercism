package phonenumber

import (
	"errors"
	"fmt"
	"strings"
)

func Number(phoneNumber string) (string, error) {
	var formatted strings.Builder
	processed := 0
	prefixProcessed := false

	for i := range len(phoneNumber) {
		char := phoneNumber[i]
		switch {
		case char == '+' || char == ' ' || char == '-' || char == '(' || char == ')' || char == '.':
			continue

		case processed == 0 || processed == 3:
			if char < '2' || char > '9' {
				// Prefix detection
				if !prefixProcessed && processed == 0 && char == '1' {
					prefixProcessed = true
					continue
				}
				return "", errors.New("Invalid digit")
			}
			formatted.WriteByte(char)
			processed++

		case char >= '0' && char <= '9':
			formatted.WriteByte(char)
			processed++

		default:
			return "", errors.New("Invalid digit")
		}
	}

	if processed == 10 {
		return formatted.String(), nil
	}
	return "", errors.New("Incorrect number length.")
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return number[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	areaCode := number[0:3]
	firstChunk := number[3:6]
	secondChunk := number[6:10]

	return fmt.Sprintf("(%s) %s-%s", areaCode, firstChunk, secondChunk), nil
}
