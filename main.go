package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Todo struct {
	Id        string `json:"_id"`
	Completed bool   `json:"completed"`
	Owner     string `json:"owner"`
	Title     string `json:"title"`
}

var todoCollection []Todo
var initialTodo = Todo{
	Id:        "abcd123",
	Completed: false,
	Owner:     "xvz123",
	Title:     "My first todo!",
}

func handler(w http.ResponseWriter, r *http.Request) {
	todoCollection = append(todoCollection, initialTodo)
	json, _ := json.MarshalIndent(todoCollection, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
