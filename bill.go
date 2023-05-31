package main

import "fmt"

// Struct
// Created a custom structure consisting of required variables
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
/*
* Function newBill
* parameters: name of type string
* return type = struct bill
 */

func newBill(name string) bill {
	// Creating first bill object
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}

	return b
}

/*
* Receiver functions
* (b bill) will only be called/received by the function of type bill
* %-25v will make the value 25 characters long and add the remaning space right of the value
* %25v will do the same and add the remaining space left to the value
 */
// format the bill
func (b bill) format() string {
	fs := "Bill breakdown \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}
