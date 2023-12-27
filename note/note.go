package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v, has the following content: \n\n%v", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ToLower(note.Title)
	fileName = strings.ReplaceAll(fileName, " ", "_")
	fileName = fmt.Sprintf("%v.json", fileName)
	
	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (Note, error) {
	if title == "" {
		return Note{}, errors.New("no title found")
	}
	
	if content == "" {
		return Note{}, errors.New("no content found")
	}

	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
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