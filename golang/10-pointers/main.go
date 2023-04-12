package main

import "fmt"

var ex *int

type Todo struct {
	id   int
	name string
}

var todos = []*Todo{}

var todo Todo = Todo{id: 1, name: "learn go"}

func main() {
	test := 1
	ex = &test
	*ex = *ex + 2
	print(*ex)

	saveTodo(&todo)
	fmt.Println(todos)
}

func saveTodo(todo *Todo) {
	todos = append(todos, todo)
}
