package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello World!")
	RunServer()
}

func RunServer() {
	router := mux.NewRouter()
	router.HandleFunc("/todos", getAllTodos).Methods("GET")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("There's an error with the server ", err)
	}

}

func getAllTodos(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")
	test := map[string]string{"name": "jeffrey"}
	err := json.NewEncoder(writer).Encode(&test)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}
}

//Create todo: POST /todos
//Retrieve all todos: GET /todos
//Retrieve todo by id: GET /todos/{id}
//Update todo by id: PUT /todos/{id}
//Delete todo by id: DELETE /todos/{id}

// Gorilla Mux
// sql library + Mysql docker container

// Tasks:
// design your project structure
// write unit tests
// setup http server
// implement handler
// (optional) basic UI using the API endpoints
// Acceptance criteria: should be able to reach each endpoint using curl or postman
