package main

import "fmt"

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// To make new bills (Create/intialze empty object of bill)

func NewBill(name string) Bill {
	b := Bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

//To format the bill using subclass or method

func (b Bill) format() string {

	fs := "Bill breakdown: \n"
	total := 0.0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v......$%v\n", k+":", v)
		total += v
	}
	//tip
	fs += fmt.Sprintf("%-25v......$%0.2f\n", "Tips:", b.tip)
	//Total
	fs += fmt.Sprintf("%-25v......$%0.2f", "Total:", total+b.tip)

	return fs
}

func (b *Bill) updateTip(tip float64) {

	b.tip = tip
}

// To add items to the bill

func (b *Bill) addItems(name string, price float64) {

	b.items[name] = price

}
