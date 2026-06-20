package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func saveTodos(todos []Todos) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("todos.json", data, 0644)
}

func storeTodos(todo Todos) (string, error) {
	var todos []Todos
	data, err := os.ReadFile("todos.json")
	if err != nil {
		if !os.IsNotExist(err) {
			return "", fmt.Errorf("failed to read file: %w", err)
		}
	} else if len(data) > 0 {
		err = json.Unmarshal(data, &todos)
		if err != nil {
			return "", fmt.Errorf("failed to parse file: %w", err)
		}
	}

	todos = append(todos, todo)
	err = saveTodos(todos)
	return "todos.json", err
}

func deleteTodo(Id string) error {
	todos, err := listtodo()
	if err != nil {
		return err
	}

	for i, t := range todos {
		if t.Id == Id {
			todos = append(todos[:i], todos[i+1:]...)
			return saveTodos(todos)
		}
	}

	return fmt.Errorf("todo with id %s not found", Id)
}

func doneTodo(Id string) error {
	todos, err := listtodo()
	if err != nil {
		return err
	}

	for i, t := range todos {
		if t.Id == Id {
			todos[i].Done = true
			return saveTodos(todos)
		}
	}

	return fmt.Errorf("todo with id %s not found", Id)
}

func listtodo() ([]Todos, error) {
	data, err := os.ReadFile("todos.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Todos{}, nil
		}
		return nil, err
	}

	if len(data) == 0 {
		return []Todos{}, nil
	}

	var todos []Todos
	err = json.Unmarshal(data, &todos)
	return todos, err
}
