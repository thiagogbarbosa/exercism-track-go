package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    percentage:= float64(successRate/100)
    successfulCarsMade := float64(productionRate)*percentage
    return successfulCarsMade
	panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    percentage:= float64(successRate/100)
    successfulCarsMade := float64(productionRate)*percentage
    successfulCarsMadePerMinute:= int(successfulCarsMade)/60
    return successfulCarsMadePerMinute
    
	panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    const tenDescount = 95000
    const individualCarCost = 10000
    groupsOfTen := carsCount / 10
	remainingCars := carsCount-(groupsOfTen*10)
    costOfProduction := (groupsOfTen * tenDescount + remainingCars*individualCarCost)
    return uint(costOfProduction)
	panic("CalculateCost not implemented")
}
