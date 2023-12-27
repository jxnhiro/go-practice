package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"todo"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	
	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(content string) (Todo, error) {	
	if content == "" {
		return Todo{}, errors.New("no content found")
	}

	return Todo{
		Text: content,
	}, nil
}

//Best practice
// func New(title, content string) (*Note, error) {
// 	if title == "" {
// 		return nil, errors.New("no title found")
// 	}
	
// 	if content == "" {
// 		return nil, errors.New("no content found")
// 	}

// 	return &Note{
// 		title: title,
// 		content: content,
// 		createdAt: time.Now(),
// 	}, nil
// }