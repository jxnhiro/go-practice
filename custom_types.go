package main

import "fmt"

//Define a custom type that also has string properties and methods
type str string

func (text str) log () {
	fmt.Println(text)
}
func main() {
	var name str = "Max"

	name.log()
}