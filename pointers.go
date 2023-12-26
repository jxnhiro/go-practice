package main

import "fmt"

func main() {
	age := 32

	var agePointer *int
	agePointer = &age

	fmt.Println("Age:", agePointer)
	adultYears := getAdultYears(age)
	fmt.Println("Adult Years:", adultYears)
}

func getAdultYears(age int) (adultYears int) {
	adultYears = age - 18
	return adultYears
}