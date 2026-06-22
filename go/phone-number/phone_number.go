package phonenumber

import (
	"errors"
	"fmt"
)

func Number(phoneNumber string) (string, error) {
	digits := make([]byte, 0, 11)
	for i := range len(phoneNumber) {
		char := phoneNumber[i]
		if char >= '0' && char <= '9' {
			digits = append(digits, char)
		} else if invalidSpecialCharacter(char) {
			return "", errors.New("Invalid character")
		}
	}

	if len(digits) == 11 {
		if digits[0] != '1' {
			return "", errors.New("Invalid country code")
		}
		digits = digits[1:]
	}

	if len(digits) != 10 {
		return "", errors.New("Incorrect length")
	}

	if digits[0] < '2' || digits[3] < '2' {
		return "", errors.New("Invalid area or exchange code")
	}

	return string(digits), nil
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

func invalidSpecialCharacter(char byte) bool {
	return char != '+' && char != ' ' && char != '-' && char != '(' && char != ')' && char != '.'
}
