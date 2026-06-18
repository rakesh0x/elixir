package main

import (
	"github.com/google/uuid"
)

type Todos struct {
	Id          string
	Objective   string
	Done        bool
	Description string
}

func createTodoStruct(Objective string, Description string) Todos {
	_id := uuid.New().String()
	TodoStruct := Todos{Id: _id, Objective: Objective, Done: false, Description: Description}
	return TodoStruct
}