package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func profit_calculator() {
	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	expenses, err := getUserInput("Expenses: ")
	
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	if expenses > revenue {
		fmt.Print("Expenses must not be more than revenue.")
		return
	}

	taxRate, err := getUserInput("Tax Rate: ")

	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	earningsBeforeTax, profit, ratio := calculateRatio(revenue, expenses, taxRate)
	err = saveCalculation(earningsBeforeTax, profit, ratio)

	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	fmt.Printf("Earnings Before Tax: %.1f\nProfit: %.1f\nRatio: %.3f\n", earningsBeforeTax, profit, ratio)
}

func saveCalculation(earningsBeforeTax, profit, ratio float32) (err error) {
	calculationText := fmt.Sprint(earningsBeforeTax, profit, ratio)
	err = os.WriteFile("profit_calculation.txt", []byte(calculationText), 0644)

	if err != nil {
		return err
	}

	return nil
}

func getUserInput(infoText string) (output float32, err error){
	fmt.Print(infoText)
	fmt.Scan(&output)

	if output < 0 {
		info := strings.Split(infoText, ":")[0]
		fmt.Printf("%v must not be less than zero.", info)

		return 0.0, errors.New("user input was less than zero")
	}

	return output, nil
}

func calculateRatio(revenue, expenses, taxRate float32) (earningsBeforeTax, profit, ratio float32){
	earningsBeforeTax = revenue - expenses
	profit = earningsBeforeTax - (earningsBeforeTax * (taxRate / 100))
	ratio = earningsBeforeTax / profit

	return earningsBeforeTax, profit, ratio
}

