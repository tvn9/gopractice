// Truct practice exmaple 1
package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {

	casey := Passenger{"Casey", 1, false}
	fmt.Println(casey)

	bill := Passenger{Name: "Bill", TicketNumber: 2}
	ella := Passenger{Name: "Ella", TicketNumber: 3}

	fmt.Println(bill, ella)

	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4
	fmt.Println(heidi)

	casey.Boarded = true
	bill.Boarded = true

	if casey.Boarded {
		fmt.Println(casey.Name, "has boarded the bus")
	}
	if bill.Boarded {
		fmt.Println(bill.Name, "has boarded the bus")
	}

	heidi.Boarded = true

	b001 := Bus{FrontSeat: heidi}

	fmt.Println(b001)
	fmt.Println(b001.FrontSeat.Name, "is in the front seat")

}
