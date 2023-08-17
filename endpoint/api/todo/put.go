package todo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"todo-application/endpoint/api/todo/queries"
	"todo-application/model"

	"github.com/gorilla/mux"
)

func UpdateHandler(queries queries.Queries) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)

		id, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Println("Invalid ID")
		}

		var todo model.Todos
		json.NewDecoder(request.Body).Decode(&todo)
		err = queries.UpdateTodoData(id, todo)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
