package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, preptime int) int {
	if preptime == 0 {
		preptime += 2
		return len(layers) * preptime
	} else {
		return len(layers) * preptime
	}
	panic("PreparationTime not implemented")
}

// TODO: define the 'Quantities()' function
// For each noodle layer in your lasagna, you will need 50 grams of noodles
// For each sauce layer in your lasagna, you will need 0.2 liters of sauce
func Quantities(layers []string) (noodles int, sauce float64) {
	noodlesLayers := 0
	sauceLayers := 0
	for _, v := range layers {
		if v == "noodles" {
			noodlesLayers += 1
		} else if v == "sauce" {
			sauceLayers += 1
		}
	}
	noodles = noodlesLayers * 50
	sauce = float64(sauceLayers) * 0.2
	return noodles, sauce
	panic("Quantities not implemented")

}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendlist, mylist []string) {
	mylist[len(mylist)-1] = friendlist[len(friendlist)-1]
	return
	panic("AddSecretIngredient not implemented")
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaledQuantities []float64
	for _, v := range quantities {
		scaledQuantities = append(scaledQuantities, v/2*float64(portions))
	}
	return scaledQuantities
	panic("ScaleRecipe not implemented")
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
