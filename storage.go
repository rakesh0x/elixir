package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//store all the todos
func storeTodos(todo Todos) (string, error) {

	filename := "todos.json"
	var err error

	//append the json to todos.json
	_, err = os.Stat(filename)
	if os.IsNotExist(err) {
			_, err = os.Create("todos.json")
		if err != nil {
			panic(err)
		}
	}

	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	//store the value in the bytes
	var todos []Todos
	if len(fileBytes) > 0 {
		err = json.Unmarshal(fileBytes, &todos)
		if err != nil {
			fmt.Println(err)
		}
	}

	data := append(todos, todo)

	dataBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filename, dataBytes, 0644) 
	if err != nil {
		fmt.Println(err)
	}

	return filename, nil
}

//delete a todo
func deleteTodo(Id string) error {
	data, err := listtodo()
	if err != nil {
		return err
	}

	for i, t := range data {
		if t.Id == Id {
			data = append(data[:i], data[i+1:]...)
			break
		}
	}

	return fmt.Errorf("todo with %s not found", Id)
}

//list all todos
func listtodo() ([]Todos, error) {
	data, err := os.ReadFile("todos.json")
	if err != nil {
		return []Todos{}, nil
	}

	var todos []Todos
	err = json.Unmarshal(data, &todos)
	return todos, err
}
