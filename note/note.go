package note

import (
	"errors"
	"time"
)

type Note struct {
	title string
	content string
	createdAt time.Time
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

func New(title, content string) (Note, error) {
	if title == "" {
		return nil, errors.New("no title found")
	}
	
	if content == "" {
		return nil, errors.New("no content found")
	}

	return Note{
		title: title,
		content: content,
		createdAt: time.Now(),
	}, nil
}