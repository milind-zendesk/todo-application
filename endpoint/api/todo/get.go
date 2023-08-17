package todo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"todo-application/endpoint/api/todo/queries"

	"github.com/gorilla/mux"
)

func GetAllHandlers(queries queries.Queries) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		todos, err := queries.GetAllTodosData()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			// _, _ = writer.Write(serialisedError)
		}

		err = json.NewEncoder(writer).Encode(&todos)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
	}
}

func GetHandlers(queries queries.Queries) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		params := mux.Vars(request)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Println("Invalid ID")
		}

		todo, err := queries.GetTodoData(id)
		if err != nil {
			panic(err)
		}

		err = json.NewEncoder(writer).Encode(&todo)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
	}
}
