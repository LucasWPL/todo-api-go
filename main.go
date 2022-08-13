package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Criar uma API em GO", Completed: false},
	{ID: "2", Item: "Aprender sobre microserviços", Completed: false},
	{ID: "3", Item: "Resolver bug no sistema", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": "todo not found"},
		)

		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*todo, error) {
	for index, todo := range todos {
		if todo.ID == id {
			return &todos[index], nil
		}
	}

	return nil, errors.New("todo not found")
}

func getTodoIndexById(id string) (int, error) {
	for index, todo := range todos {
		if todo.ID == id {
			return index, nil
		}
	}

	return 0, errors.New("todo not found")
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": "todo not found"},
		)

		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func deleteTodo(context *gin.Context) {
	id := context.Param("id")
	index, err := getTodoIndexById(id)

	if err != nil {
		context.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": "todo not found"},
		)

		return
	}

	todos = append(todos[:index], todos[index+1:]...)

	context.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", addTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.DELETE("/todos/:id", deleteTodo)

	router.Run("localhost:9090")
}
