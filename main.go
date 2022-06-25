package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Ramayan", Completed: false},
	{ID: "2", Item: "Mahabharat", Completed: false},
	{ID: "3", Item: "Kuran", Completed: false},
}

//--------------------------<< get all todo >>---------------
func getTodos(context *gin.Context) { //context is info of header, data etc
	context.IndentedJSON(http.StatusOK, todos) // prepare response body(json) from struct

}

//--------------------------<< add todo >>---------------
func addTodos(context *gin.Context) {
	//need to convert data comming from post api in struct(json)
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, todos)
}

//--------------------------<< Find todo By Id >>---------------
func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func findTodo(context *gin.Context) {
	// localhost:9099/todos/3 need to parse param 3
	// all info is available in context
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

//--------------------------<< update todo >>---------------
func toggleStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	fmt.Println("Running...")
	//Creating a gin server
	router := gin.Default()
	//Create an end point which return us todos list
	router.GET("/todos", getTodos)
	//Create an end point which add todo in my todos list
	router.POST("/todos", addTodos)
	//adding dynamic end point
	router.GET("/todos/:id", findTodo)
	//update todo
	router.PATCH("/todos/:id", toggleStatus)
	//Specifying router path
	router.Run("localhost:9099")

}
