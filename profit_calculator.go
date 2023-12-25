package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxRate float32

	print("Revenue: ")
	fmt.Scan(&revenue)

	print("Expenses: ")
	fmt.Scan(&expenses)

	print("Tax Rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax, profit, ratio := calculateRatio(revenue, expenses, taxRate)

	fmt.Print(earningsBeforeTax, profit, ratio)
}

func print(text string) {
	fmt.Print(text)
}

func calculateRatio(revenue, expenses, taxRate float32) (earningsBeforeTax, profit, ratio float32){
	earningsBeforeTax = revenue - expenses
	profit = earningsBeforeTax - (earningsBeforeTax * (taxRate / 100))
	ratio = earningsBeforeTax / profit

	return earningsBeforeTax, profit, ratio
}

