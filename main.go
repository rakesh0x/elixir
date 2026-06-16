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
		createTodoStruct(Objective, args[0])
		fmt.Println("creating a todo")
	case "list":
		createTodoStruct(Objective, args[1])
		fmt.Println("listing all the todo")
	case "delete": 
		createTodoStruct(Objective, args[2])
		fmt.Println("deleting a todo")
	case "done":
		createTodoStruct(Objective, args[3])
		fmt.Println("completed a todo")

	default: 
		fmt.Println("Unknown command")
	}
}