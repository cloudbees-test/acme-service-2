package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Todo struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

var (
	todos    []Todo
	nextID   = 1
	todoLock sync.Mutex
)

func main() {
	appendTodo(Todo{Text: "Buy milk", Completed: false})
	appendTodo(Todo{Text: "Take out garbage", Completed: false})

	http.HandleFunc("/", handleTodos)
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getTodos(w)
	case "POST":
		createTodo(w, r)
	case "PUT":
		updateTodo(w, r)
	case "DELETE":
		deleteTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTodos(w http.ResponseWriter) {
	todoLock.Lock()
	defer todoLock.Unlock()

	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	appendTodo(todo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	// ... Implementation for updating a todo
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	// ... Implementation for deleting a todo
}

func appendTodo(todo Todo) {
	todoLock.Lock()
	defer todoLock.Unlock()

	todo.ID = nextID
	nextID++
	todos = append(todos, todo)
}
