package main

import "fmt"

func variadic_functions() {
	fmt.Println(sumUp(1,2,3,4,5))
}

func sumUp(integers ...int) int {
	sum := 0
	for _, value := range integers {
		sum += value
	}
	return sum
}