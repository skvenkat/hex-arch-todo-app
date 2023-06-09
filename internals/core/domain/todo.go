package domain

import "fmt"

type ToDo struct {
	ID          string
	Title       string
	Description string
}

func NewToDo(id, title, description string) *ToDo {
	return &ToDo{
		ID:          id,
		Title:       title,
		Description: description,
	}
}

func (t *ToDo) String() string {
	return fmt.Sprintf("%s - %s", t.Title, t.Description)
}
