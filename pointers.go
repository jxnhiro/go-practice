package main

import "fmt"

func main() {
	age := 32
	var agePointer *int = &age

	fmt.Println(agePointer)
}