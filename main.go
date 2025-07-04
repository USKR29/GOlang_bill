package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	switch opt {
	case "a":
		name, _ := getInput("Item name:", reader)
		price, _ := getInput("Item price:", reader)
		fmt.Println(name, price)

		//convert string to float of price

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be number")
			options(b)
		}
		b.addItems(name, p)
		fmt.Println("Item has been added sucessfully", name, price)
		options(b)
	case "t":
		tip, _ := getInput("Enter the tip amount $:", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Tip must be only numbers")
			options(b)
		}
		b.updateTip(t)
		fmt.Println("you have added tip:", tip)
		options(b)
	case "s":
		fmt.Println("You chose to save the bill", b.format())

	default:
		fmt.Println("Please enter valid key word")
		options(b)

	}
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
