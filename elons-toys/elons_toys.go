package elon

import "fmt"

func (myCar *Car) Drive() {
	if myCar.battery >= myCar.batteryDrain {
		(*myCar).battery -= myCar.batteryDrain
		(*myCar).distance += myCar.speed
	}

}

func (myCar *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", (*myCar).distance)
}

func (myCar *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", (*myCar).battery)
}

func (myCar *Car) CanFinish(trackDistance int) bool {
	// Calculate how many times the car can move before the battery is depleted.
	maxMovements := myCar.battery / myCar.batteryDrain
	// Calculate the total distance the car can cover with the available battery.
	totalDistance := maxMovements * myCar.speed

	// If the total distance is greater than or equal to the track distance, the car can finish.
	return totalDistance >= trackDistance
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
