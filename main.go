package main

import (
	"os"
	"fmt"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Please provide a command")
		return
	}

	switch args[1] {
	case "create":
		if len(args) < 3 {
			fmt.Println("Usage: todo create <objective> [description]")
			return
		}
		description := ""
		if len(args) > 3 {
			description = args[3]
		}
		todo := createTodoStruct(args[2], description)
		storeTodos(todo)
		fmt.Println("Created todo:", todo.Id)

	case "list":
		todos, err := listtodo()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if len(todos) == 0 {
			fmt.Println("No todos found")
			return
		}
		for _, t := range todos {
			status := "[ ]"
			if t.Done {
				status = "[x]"
			}
			fmt.Printf("%s %s | %s | %s\n", status, t.Id[:8], t.Objective, t.Description)
		}

	case "delete":
		if len(args) < 3 {
			fmt.Println("Usage: todo delete <id>")
			return
		}
		err := deleteTodo(args[2])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Deleted todo:", args[2])

	case "done":
		if len(args) < 3 {
			fmt.Println("Usage: todo done <id>")
			return
		}
		err := doneTodo(args[2])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Marked done:", args[2])

	default:
		fmt.Println("Unknown command:", args[1])
	}
}
