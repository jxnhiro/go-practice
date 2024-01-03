package main

import "fmt"

func main() {
	age := 32

	mutateAdultYears(&age)
	fmt.Printf("Adult Years: %v", age)
}

func getAdultYears(age *int) int {
	return *age - 18
}

//Dereference to mutate in-memory
func mutateAdultYears(age *int) {
	*age = *age - 18
}