package main

import "fmt"

type floatMap map[string]float64 

func make_function() {
	foo := make([]string, 2, 4)
	foo[0] = "Hello"
	foo[1] = "Hi"

	fmt.Println(foo)

	for index, value := range foo {
		fmt.Println(index, value)
	}

	courseRatings := make(map[string]float64, 2)

	for key, value := range courseRatings {
		fmt.Println(key, value)
	}
}