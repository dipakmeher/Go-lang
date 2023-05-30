package main

// Struct
// Created a custom structure consisting of required variables
type bill struct {
	name string
	item map[string]float64
	tip  float64
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
		name: name,
		item: map[string]float64{},
		tip:  0,
	}

	return b
}
