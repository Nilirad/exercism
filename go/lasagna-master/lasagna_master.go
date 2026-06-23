package lasagnamaster

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}

	return len(layers) * time
}

func Quantities(layers []string) (int, float64) {
	noodleLayers := 0
	sauceLayers := 0

	for i := range len(layers) {
		switch layers[i] {
		case "noodles":
			noodleLayers++
		case "sauce":
			sauceLayers++
		}
	}

	return 50 * noodleLayers, 0.2 * float64(sauceLayers)
}

func AddSecretIngredient(friendsList []string, myList []string) {
	secret := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = secret
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaled := make([]float64, 0, len(quantities))
	for i := range len(quantities) {
		scaled = append(scaled, quantities[i]*float64(portions)/2.0)
	}

	return scaled
}
