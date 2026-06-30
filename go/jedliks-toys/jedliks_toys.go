package jedlik

import "fmt"

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func (car *Car) Drive() {
	finalBattery := car.battery - car.batteryDrain
	if finalBattery < 0 {
		return
	}

	car.distance += car.speed
	car.battery = finalBattery
}

func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// CanFinish checks if a car is able to finish a certain track.
func (car Car) CanFinish(distance int) bool {
	drain := distance * car.batteryDrain / car.speed
	finalBattery := car.battery - drain
	return finalBattery >= 0
}
