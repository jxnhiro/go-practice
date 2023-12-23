package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxRate float32

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	// Calculate earnings before tax
	earningsBeforeTax := revenue - expenses

	// Calculate earnings after tax
	profit := earningsBeforeTax - (earningsBeforeTax * (taxRate / 100))

	// Calculate ratio
	ratio := earningsBeforeTax / profit

	fmt.Print(earningsBeforeTax, profit, ratio)
}