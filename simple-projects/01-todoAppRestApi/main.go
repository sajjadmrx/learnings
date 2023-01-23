package main

import (
	"e/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/todos", handlers.GetTodos)
	router.POST("/todos", handlers.AddTodo)
	router.GET("/todos/:todoId", handlers.GetTodoById)
	router.PUT("/todos/:todoId/complete", handlers.ToggleTodoComplete)

	router.Run("localhost:9090")
	
}
