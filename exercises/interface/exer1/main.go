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

package main

import "fmt"

type Director interface {
	LiftDirector()
}

const (
	largeL    = "Large Lift"
	standardL = "Standard Lift"
	smallL    = "Small Lift"
)

type Vehicles struct {
	vType string
	model string
}

func (v *Vehicles) LiftDirector() {
	switch v.vType {
	case "Car":
		fmt.Printf("Send %s %s to %s\n", v.vType, v.model, standardL)
	case "Truck":
		fmt.Printf("Send %s %s to %s\n", v.vType, v.model, largeL)
	case "Motorcycles":
		fmt.Printf("Send %s %s to %s\n", v.vType, v.model, smallL)
	default:
		fmt.Printf("Un-supporte vihicle type")
	}
	fmt.Println()
}

func LiftDirector(vihicles []Vehicles) {
	for _, t := range vihicles {
		t.LiftDirector()
	}
}

func main() {
	t1 := Vehicles{"Truck", "F150"}
	c1 := Vehicles{"Car", "Prius 2020"}
	m1 := Vehicles{"Motorcycles", "Magna 650"}

	t1.LiftDirector()
	c1.LiftDirector()
	m1.LiftDirector()

	vehicles := []Vehicles{t1, c1, m1}
	LiftDirector(vehicles)
}
