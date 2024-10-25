package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var rate float64 = float64(productionRate) * (successRate / 100)
	return rate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {

	return int((float64(productionRate) * (successRate) / 100) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupedCar := int(carsCount / 10)
	indivialCar := carsCount % 10

	return uint(groupedCar*95000 + indivialCar*10000)

}
