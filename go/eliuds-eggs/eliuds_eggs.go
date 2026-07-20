package eliudseggs

const mask = 1

func EggCount(displayValue int) int {
	eggs := 0
	for displayValue != 0 {
		displayValue = displayValue & (displayValue - 1)
		eggs++
	}

	return eggs
}
