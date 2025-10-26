package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type ToDo struct {
	Content   string    `json:"content"`
}

func (toDo ToDo) Display() {
fmt.Println(toDo)
}

func (toDo ToDo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(toDo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(content string) (ToDo, error) {
	if content == "" {
		return ToDo{}, errors.New("invalid input")
	}

	return ToDo{
		Content:   content,
	}, nil
}