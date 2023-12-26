package main

import "fmt"

func main() {
	age := 32

	var agePointer *int = &age

	fmt.Println("Age:", *agePointer)
	editAge(agePointer)
	fmt.Println("Adult Years:", age)
}

//No copy of 32 made
func editAge(age *int) {
	*age = *age - 18
}