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

	fmt.Println(args)

	switch args[1] {
	case "create":
		fmt.Println("creating a todo")
	case "list":
		fmt.Println("listing all the todo")
	case "delete": 
		fmt.Println("deleting a todo")
	case "done":
		fmt.Println("completed a todo")

	default: 
		fmt.Println("Unknown command")
	} 

}
