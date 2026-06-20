# todo-go

a simple cli todo app written in go.

## usage

```
go run . create "<objective>" [description]
go run . list
go run . done <id>
go run . delete <id>
```

## commands

**create** — add a new todo. pass the task as the first argument and an optional description as the second.

**list** — show all todos. completed ones have a [x] and pending ones have a [ ].

**done** — mark a todo as done using its id.

**delete** — remove a todo using its id.

## how it works

all todos are stored in a `todos.json` file in the same directory. each todo has an id (uuid), objective, description, and a done flag.

## what i learned

- go slices are way more flexible than arrays
- using `os.args` to parse cli input
- reading and writing json files with `encoding/json`
- the uuid package from google
