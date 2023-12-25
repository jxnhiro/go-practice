package main

import (
	"fmt"
)

func main() {
	var revenue float32 = getUserInput("Revenue: ")
	var expenses float32 = getUserInput("Expenses: ")
	var taxRate float32 = getUserInput("Tax Rate: ")

	earningsBeforeTax, profit, ratio := calculateRatio(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax: %.1f\nProfit: %.1f\nRatio: %.3f\n", earningsBeforeTax, profit, ratio)
}

func getUserInput(infoText string) (output float32){
	fmt.Print(infoText)
	fmt.Scan(&output)
	return output
}

func calculateRatio(revenue, expenses, taxRate float32) (earningsBeforeTax, profit, ratio float32){
	earningsBeforeTax = revenue - expenses
	profit = earningsBeforeTax - (earningsBeforeTax * (taxRate / 100))
	ratio = earningsBeforeTax / profit

	return earningsBeforeTax, profit, ratio
}

