package todo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"todo-application/database"
	"todo-application/endpoint/api/todo/queries"

	"github.com/gorilla/mux"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	var con, _ = database.Connection()
	writer.Header().Set("Content-Type", "application/json")

	todos, err := queries.GetAllTodosData(con)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(writer).Encode(&todos)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}
}

func Get(writer http.ResponseWriter, request *http.Request) {
	var con, _ = database.Connection()
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Invalid ID")
	}

	todo, err := queries.GetTodoData(con, id)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(writer).Encode(&todo)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}
}
