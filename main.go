package main

import "fmt"

func main() {
	//newBill is the function in bill.go file
	mybill := newBill("mario's bill")

	mybill.updateTip(10)

	mybill.addItem("paneer", 2.99)

	fmt.Println(mybill.format())

}
