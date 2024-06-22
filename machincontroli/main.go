package main

import "fmt"

type Car struct {
	speed   int
	battery int
}

func NewCar(speed, battery int) *Car {
	var ncar = Car{speed: speed, battery: battery}
	return &ncar
}
func GetSpeed(car *Car) int {
	var gspeed = &car.speed
	return *gspeed
}
func GetBattery(car *Car) int {
	var gbattery = &car.battery
	return *gbattery
}
func ChargeCar(car *Car, minutes int) {
	car.battery = car.battery + minutes/2
	if car.battery > 100 {
		car.battery = 100
	}
}
func TryFinish(car *Car, distance int) string {
	var time = float32(distance) / float32(car.speed)
	if distance/2 > car.battery {
		car.battery = 0
		return ""
	} else {
		car.battery = car.battery - distance/2
		res := fmt.Sprintf("%f", time)
		return res
	}
}
