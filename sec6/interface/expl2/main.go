// Interface example
package main

import "fmt"

type Preparer interface {
	PrepareDish()
	PrintDisk()
}

type Salad string
type Dressing string

func (s Salad) PrepareDish() {
	fmt.Println("Chop salad")
}

func (d Dressing) PrepareDish() {
	fmt.Println("Mix dressing")
}

func PrepareDish(dishes []Preparer) {
	fmt.Println("Prepare diskes:")
	for i, d := range dishes {
		fmt.Printf("#%d disk %s -- ", i+1, d)
		d.PrepareDish()
	}
	fmt.Println()
}

func (s Salad) PrintDisk() {
	fmt.Println(s)
}

func (d Dressing) PrintDisk() {
	fmt.Println(d)
}

func main() {
	s1 := Salad("New York Salad")
	s2 := Salad("Italian Salad")
	d1 := Dressing("Italian Dressing")
	d2 := Dressing("Japan Dressing")

	dishes := []Preparer{s1, s2, d1, d2}
	PrepareDish(dishes)
}
