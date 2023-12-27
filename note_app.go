package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jxnhiro/go-practice/note"
	"github.com/jxnhiro/go-practice/todo"
)

func main() {
	title, content := getNoteData()
	todoText := getTodoData()

	todo, todoErr := todo.New(todoText)
	
	if todoErr != nil {
		fmt.Println(todoErr)
		return
	}
	
	note, noteErr := note.New(title, content)

	if noteErr != nil {
		fmt.Println(noteErr)
		return
	}

	todo.Display()
	todoErr = todo.Save()

	if todoErr != nil {
		fmt.Println("Saving the TODO failed.")
		return
	}

	fmt.Println("Saving the TODO succeeded.")

	note.Display()
	noteErr = note.Save()

	if noteErr != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note succeeded.")
}

func getTodoData() (string) {
	return getUserInput("Todo Text: ")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) (text string){
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	//In windows, line breaks end with \n\r, so we also delete \r
	text = strings.TrimSuffix(text, "\r")

	return text
}