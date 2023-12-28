package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jxnhiro/go-practice/note"
	"github.com/jxnhiro/go-practice/todo"
)

type saver interface {
	Save() error 
}

//Embedded type
type outputtable interface {
	saver
	Display()
}

//Alternative syntax
// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	printSomething(1.0)
	printSomething(false)

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

	noteOutputErr := outputData(note)

	if noteOutputErr != nil {
		return
	}

	outputData(todo)
}

//Any value or any type is allowed in Go with interface{}
func printSomething(value interface{}) {
	//Add type switches
	switch value.(type){
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float: ", value)
	case string:
		fmt.Println("String: ", value)
	}
}
func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the data failed.")
		return err
	}

	fmt.Println("Saving the data succeeded.")
	return nil
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