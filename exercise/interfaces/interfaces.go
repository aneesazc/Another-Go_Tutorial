//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

// Code here:
package main

import "fmt"

type Vehicle struct {
	Type  string
	Model string
}

type LiftVehicle interface {
	vehicleLift()
}

func (v Vehicle) vehicleLift() {
	switch v.Type {
	case "Motorcycle":
		fmt.Printf("Motorcycle %s is on a small lift\n", v.Model)
	case "Car":
		fmt.Printf("Car %s is on a standard lift\n", v.Model)
	case "Truck":
		fmt.Printf("Truck %s is on a large lift\n", v.Model)
	}
}

func main() {
	motorcycle := Vehicle{Type: "Motorcycle", Model: "Harley"}
	car := Vehicle{Type: "Car", Model: "Corolla"}
	truck := Vehicle{Type: "Truck", Model: "F150"}

	vehicles := []LiftVehicle{motorcycle, car, truck}
	fmt.Println("Vehicles at the shop:", vehicles)

	for _, vehicle := range vehicles {
		vehicle.vehicleLift()
	}
}
