package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content, error := getNoteData()

	if error != nil {
		fmt.Println(error)
		return
	}
}

func getNoteData() (string, string, error) {
	title, error := getUserInput("Note title:")

	if error != nil {
		fmt.Println(error)
		return "", "", error
	}

	content, error := getUserInput("Note content:")

	if error != nil {
		fmt.Println(error)
		return "", "", error
	}

	return title, content, nil
}

func getUserInput(prompt string) (value string, err error){
	fmt.Print(prompt)
	fmt.Scanln(&value)
	
	if value == "" {
		return "", errors.New("Invalid input.")
	}

	return value, nil
}