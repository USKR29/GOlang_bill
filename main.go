package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// for resuable input reader
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, error := r.ReadString('\n')

	return strings.TrimSpace(input), error

}

// Getting the bill details from the user
func getBill() Bill {

	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Enter the name of the bill:")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Enter the name of the bill:", reader)

	b := NewBill(name)
	fmt.Println("Created the bill for: ", b.name)
	return b
}

func options(b Bill) {

	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Enter your choice (a - add item, S- save the bill, T- To add tip)", reader)
	fmt.Println(opt)
}

func main() {

	// MyBill := NewBill("Jhone 's Bill")
	// MyBill.addItems("Burger", 5.69)
	// MyBill.addItems("Cake", 4.19)
	// MyBill.addItems("Juice", 5.00)
	// MyBill.addItems("Sandwich", 6.89)

	// MyBill.updateTip(10)
	// fmt.Println(MyBill.format())

	mybill := getBill()

	options(mybill)

	//fmt.Println(mybill)
}
