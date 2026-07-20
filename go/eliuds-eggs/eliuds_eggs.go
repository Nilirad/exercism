package eliudseggs

const mask = 1

func EggCount(displayValue int) int {
	unsignedValue := uint(displayValue)
	eggs := 0
	for unsignedValue != 0 {
		if unsignedValue&mask == 1 {
			eggs++
		}
		unsignedValue >>= 1
	}

	return eggs
}
