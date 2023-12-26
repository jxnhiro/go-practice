package main

import "fmt"

func main() {
	age := 32

	var agePointer *int = &age

	fmt.Println("Age:", *agePointer)
	adultYears := getAdultYears(agePointer)
	fmt.Println("Adult Years:", adultYears)
}

//No copy of 32 made
func getAdultYears(age *int) (adultYears int) {
	adultYears = *age - 18
	return adultYears
}