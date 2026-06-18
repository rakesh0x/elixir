package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// load all the todos and do the operations
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

	fileBytes, err := ioutil.ReadFile(filename)
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

func deleteTodo(Id string) {
	if Id != "" {
		deletedId := Id
	}

	// do this after storing the struct array in the todos.json
}

func listtodo() {}
