package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jxnhiro/go-practice/note"
)

func main() {
	title, content := getNoteData()

	userNote, error := note.New(title, content)

	if error != nil {
		fmt.Println(error)
		return
	}

	
	userNote.Display()

	err := userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note succeeded.")
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