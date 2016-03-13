package main

import "fmt"

func describe(coffe Beverage) {
	fmt.Printf("%s $%.2f\n", coffe.GetDescription(), coffe.Cost())
}

func main() {
	espresso := NewEspresso()
	describe(espresso)

	dr := NewDarkRoast()
	dr = NewMocha(dr)
	dr = NewMocha(dr)
	dr = NewWhip(dr)
	describe(dr)

	hb := NewHouseBlend()
	hb = NewSoy(hb)
	hb = NewMocha(hb)
	hb = NewWhip(hb)
	describe(hb)
}
