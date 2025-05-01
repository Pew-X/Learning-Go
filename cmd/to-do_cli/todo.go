package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Done        bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

//slice to hold all todos

type Todos []Todo

//ADD

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Done:        false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
}

//validate index

func (todos *Todos) validateIndex(index int) error {

	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err.Error())
		return err
	}
	return nil

}

func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil
}

//toggle

func (todos *Todos) toggle(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Done

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	t[index].Done = !isCompleted
	return nil
}

//Edit Method

func (todos *Todos) edit(index int, title string) error {

	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title

	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Done {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)

			}

		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)

	}

	table.Render()
}
