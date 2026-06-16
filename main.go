package main

import (
	"flag"
	"fmt"
)


func main() {
	create := flag.String("create", "", "Create a todo")
	list := flag.Bool("list", false, "List all todo")
	deleteID := flag.Int("delete", 0, "Delete a todo")
	done := flag.Int("done", 0, "Completed a todo")

	fmt.Println(create)

	flag.Parse()

	if *create != "" {
		fmt.Println("Creating a todo")
	} 

	if *list {
		fmt.Println("listing all the todo")
	}

	if *deleteID != 0 {
		fmt.Println("delete a todo")
	} 

	if *done != 0{
		fmt.Println("done with a todo")
	}

}