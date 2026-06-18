package main

import (
	"os"
	"fmt"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Please provide a command")
	}

	Objective := args[1]

	fmt.Println(args)

	switch args[1] {
	case "create":
		todo := createTodoStruct(Objective, args[0])
		storeTodos(todo)
		fmt.Println("creating a todo")
	case "list":
		todo := createTodoStruct(Objective, args[1])
		storeTodos(todo)
		fmt.Println("listing all the todo")
	case "delete": 
		todo := createTodoStruct(Objective, args[2])
		storeTodos(todo)
		fmt.Println("deleting a todo")
	case "done":
		todo := createTodoStruct(Objective, args[3])
		storeTodos(todo)
		fmt.Println("completed a todo")

	default: 
		fmt.Println("Unknown command")
	}
}