package internal

import "fmt"

type Todo struct {
	Description string `json:"description"`
}


var todos []Todo

func ListTODOs() []Todo {
	return todos
}

func SaveTODO(todo *Todo) bool {
	todos = append(todos, *todo)
	fmt.Println(todos)
	return true
}

func DeleteTODO(index int) bool {
	todos = append(todos[:index], todos[index+1:]...)
	return true
}