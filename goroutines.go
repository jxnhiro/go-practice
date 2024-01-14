package main

import (
	"fmt"
	"time"
)

func greet(phrase string, done chan bool) {
	fmt.Println(phrase)
}

func slowGreet(phrase string, done chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(phrase)
	close(done)
}

func main() {
	done := make(chan bool)
	go greet("Hello", done)
	go greet("Hello", done)
	go greet("Hello", done)
	go slowGreet("H.. e.. l.. l.. o!", done)
	
	for range done {

	}
}

