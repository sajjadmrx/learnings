package handlers

import (
	"e/store"
	"e/types"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, store.Todos)
}

func GetTodoById(context *gin.Context) {
	todoId := context.Param("todoId")
	todo, err := findTodoById(todoId)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "❌ todo not found."})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func AddTodo(ctx *gin.Context) {
	var newTodo types.Todo
	err := ctx.BindJSON(&newTodo)
	if err != nil {
		log.Println(err)
		ctx.IndentedJSON(http.StatusBadRequest, "body is empty")
		return
	}

	store.Todos = append(store.Todos, newTodo)
	ctx.IndentedJSON(http.StatusOK, newTodo)
}

func ToggleTodoComplete(ctx *gin.Context) {
	todoId := ctx.Param("todoId")
	todo, err := findTodoById(todoId)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "❌ todo not found."})
		return
	}
	todo.Completed = !todo.Completed

	ctx.IndentedJSON(http.StatusOK, todo)

}

func findTodoById(todoId string) (*types.Todo, error) {
	for i, t := range store.Todos {
		if t.Id == todoId {
			return &store.Todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}
