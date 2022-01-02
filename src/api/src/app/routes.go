package app

import "todo-api/controllers"

func routes() {
	router.GET("/todos/:todo_id", controllers.GetTodo)
	router.GET("/todos", controllers.GetAllTodos)
	router.POST("/todos", controllers.CreateTodo)
}