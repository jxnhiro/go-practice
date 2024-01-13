package main

import "fmt"

type customFunctionType func(int) []int 

func main() {
	numbers := []int{1,2,3,4}
	double_number(&numbers)
	fmt.Println(numbers)
}

func double_number(numbers *[]int) {
	for index, value := range *numbers {
		(*numbers)[index] = value * 2
	}
}