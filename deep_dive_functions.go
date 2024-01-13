package main

import "fmt"

type customFunctionType func(int) []int 

func deep_dive_functions() {
	numbers := []int{1,2,3,4}
	double_number(&numbers)
	fmt.Println(numbers)

	anonymous_function_example(2, func(number int) int {
		return number * 2
	})

	doubler := factor_multiplier(2)
	fmt.Println(doubler(3))
}

func factor_multiplier(factor int) func(int) int {
	return func(number int) int {
		return factor * number
	}
}

func double_number(numbers *[]int) {
	for index, value := range *numbers {
		(*numbers)[index] = value * 2
	}
}

func anonymous_function_example(number int, transformerFunc func(int) int) int {
	return transformerFunc(number)
}