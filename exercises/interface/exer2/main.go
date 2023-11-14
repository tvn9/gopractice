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

const (
	SmallLift Lift = iota
	StandardLift
	LargeLift
)

type Lift int

type LiftPicker interface {
	PickLift() Lift
}

type Motorcycle string
type Car string
type Truck string

func (m Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle %v", string(m))
}

func (c Car) String() string {
	return fmt.Sprintf("Car %v", string(c))
}

func (t Truck) String() string {
	return fmt.Sprintf("Truck %v", string(t))
}

func (c Car) PickLift() Lift {
	return StandardLift
}

func (m Motorcycle) PickLift() Lift {
	return SmallLift
}

func (t Truck) PickLift() Lift {
	return LargeLift
}

func liftRouter(l LiftPicker) {
	switch l.PickLift() {
	case SmallLift:
		fmt.Printf("send %v to small lift\n", l)
	case StandardLift:
		fmt.Printf("send %v to standard lift\n", l)
	case LargeLift:
		fmt.Printf("send %v to large lift\n", l)
	default:
		fmt.Printf("Vehicle type not supported.")
	}
}

func main() {
	car := Car("Covette Z9")
	truck := Truck("Ford Raptor")
	motorcycle := Motorcycle("Honda VFR 800")

	liftRouter(car)
	liftRouter(truck)
	liftRouter(motorcycle)

}
