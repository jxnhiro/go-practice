package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (balance float64, err error){
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000.0, errors.New("failed to read file")
	}

	balanceText := string(data)
	balance, _ = strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000.0, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func bank() {
	var accountBalance, err = getBalanceFromFile()
	
	if err != nil {
		fmt.Println("ERROR: ", err)
		panic("Error found in application, stopping application")
	}

	for {
		fmt.Println("Welcome to Go Bank")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance?")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int 

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid input, must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated! New amount:", accountBalance)
		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid input, must be greater than 0.")
				continue
			} else if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount, you can't withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("Balcne updated! New amount:", accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for the visit!")
			return
		}


		// Using boolean as a control structure
		// checkBalance := choice == 1
		// if checkBalance {
		// 	fmt.Println("Your balance is:", accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("Your deposit: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)

		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid input, must be greater than 0.")
		// 		continue
		// 	}

		// 	accountBalance += depositAmount
		// 	fmt.Println("Balance updated! New amount:", accountBalance)
		// } else if choice == 3 {
		// 	fmt.Print("Withdrawal amount: ")
		// 	var withdrawAmount float64
		// 	fmt.Scan(&withdrawAmount)

		// 	if withdrawAmount <= 0 {
		// 		fmt.Println("Invalid input, must be greater than 0.")
		// 		continue
		// 	} else if withdrawAmount > accountBalance {
		// 		fmt.Println("Invalid amount, you can't withdraw more than you have.")
		// 		continue
		// 	}

		// 	accountBalance -= withdrawAmount
		// 	fmt.Println("Balcne updated! New amount:", accountBalance)
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	break
		// }
	}
}