package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLost struct {
	spaces []Space
}

func occupiedSpace(lot *ParkingLost, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLost) occupiedSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLost) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	lot := ParkingLost{spaces: make([]Space, 10)}
	fmt.Println("Initial:", lot)
	lot.occupiedSpace(1)
	occupiedSpace(&lot, 2)
	fmt.Println("After occupied:", lot)

	lot.vacateSpace(2)
	fmt.Println("After vacate:", lot)

}
