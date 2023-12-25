package main

import "fmt"

func main() {
	var accountBalance = 1000.0
	
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

		//Using boolean as a control structure
		checkBalance := choice == 1
		if checkBalance {
			fmt.Println("Your balance is:", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid input, must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		} else if choice == 3 {
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
			fmt.Println("Balcne updated! New amount:", accountBalance)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}

	fmt.Println("Thank you for the visit!")
}