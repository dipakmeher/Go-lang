package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Creating getInput to simplify input functionality from user
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

/*
* Bufio is used to take input from the user
* Reader is going to read the string untill user presses enter
 */
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")

	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)

	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill,  t - add tip): ", reader)
	fmt.Println(opt)
}

func main() {
	//newBill is the function in bill.go file
	mybill := createBill()

	promptOptions(mybill)

	// mybill.updateTip(10)

	// mybill.addItem("paneer", 2.99)

	// fmt.Println(mybill.format())

}
