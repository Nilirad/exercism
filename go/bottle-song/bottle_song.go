package bottlesong

import "fmt"

const linesPerVerse = 4
const constantLine = "And if one green bottle should accidentally fall,"

func integerToWordCapitalized(n int) string {
	switch n {
	case 1:
		return "One"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	case 10:
		return "Ten"
	default:
		panic("Invalid number")
	}
}

func integerToWordLowercase(n int) string {
	switch n {
	case 0:
		return "no"
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	case 10:
		return "ten"
	default:
		panic("Invalid number")
	}
}

func bottlesToString(n int) string {
	if n == 1 {
		return "bottle"
	}
	return "bottles"
}

func Recite(startBottles, takeDown int) []string {
	lines := make([]string, 0, linesPerVerse*takeDown)

	firstPass := true
	for takeDown > 0 {
		if !firstPass {
			lines = append(lines, "")
		} else {
			firstPass = false
		}
		openingLine := fmt.Sprintf(
			"%s green %s hanging on the wall,",
			integerToWordCapitalized(startBottles),
			bottlesToString(startBottles))
		startBottles--
		lastLine := fmt.Sprintf(
			"There'll be %s green %s hanging on the wall.",
			integerToWordLowercase(startBottles),
			bottlesToString(startBottles))
		takeDown--
		lines = append(lines, openingLine, openingLine, constantLine, lastLine)
	}

	return lines
}
