package main

//Generics, you can have two placeholder generic types and return both of them
func add[T int, K string](a, b T) K {
	return "Hello"
}
