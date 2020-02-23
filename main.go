package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	Id        int    `json:"_id"`
	Completed bool   `json:"completed"`
	Owner     string `json:"owner"`
	Title     string `json:"title"`
}
type PostBody struct {
	Title string
}
type PatchBody struct {
	Completed bool
}
type Error struct {
	Message string `json:"message"`
}

var i = 0
var initialTodo = Todo{
	Id:        0,
	Completed: false,
	Owner:     "xyz123",
	Title:     "My first todo!",
}
var todoCollection = []Todo{
	initialTodo,
}

func handler(w http.ResponseWriter, r *http.Request) {

	var jsonBody []uint8
	todoPath := r.URL.Path[len("/todos/"):]
	verb := r.Method

	fmt.Printf("%v %v\n", verb, r.URL.Path)

	// Respond with collection
	if todoPath == "" && verb == "GET" {

		jsonBody, _ = json.MarshalIndent(todoCollection, "", "  ")
		w.Header().Set("Content-Type", "application/json")

		// Add todo to collection
	} else if todoPath == "" && verb == "POST" {

		var newTodo PostBody
		body := json.NewDecoder(r.Body)
		err := body.Decode(&newTodo)

		if err != nil {
			jsonBody, _ = json.Marshal(Error{
				Message: "Error in decoding body",
			})
		} else {
			i++
			todoCollection = append(todoCollection, Todo{
				Id:        i,
				Completed: false,
				Owner:     "xyz123",
				Title:     newTodo.Title,
			})
		}

		// Respond with a todo from collection
	} else if todoPath != "" && verb == "GET" {

		id, _ := strconv.Atoi(todoPath)
		for _, element := range todoCollection {
			if element.Id == id {
				jsonBody, _ = json.MarshalIndent(element, "", "  ")
				break
			}
		}

		// Update a todo in collection
	} else if todoPath != "" && verb == "PATCH" {

		var updated PatchBody
		body := json.NewDecoder(r.Body)
		err := body.Decode(&updated)

		if err != nil {
			jsonBody, _ = json.Marshal(Error{
				Message: "Error in decoding body",
			})
		} else {
			id, _ := strconv.Atoi(todoPath)
			for i, element := range todoCollection {
				if element.Id == id {
					todoCollection[i].Completed = updated.Completed
					break
				}
			}
		}

		// Remove a todo from collection
	} else if todoPath != "" && verb == "DELETE" {

		var location int
		id, _ := strconv.Atoi(todoPath)
		for i, element := range todoCollection {
			if element.Id == id {
				location = i
				break
			}
		}
		todoCollection = append(todoCollection[:location], todoCollection[location+1:]...)

		// Didn't understand request; respond with empty body
	} else {
		jsonBody, _ = json.Marshal([]Todo(nil))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBody)
}

func main() {
	http.HandleFunc("/todos/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
