package todo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo-application/database"
	"todo-application/endpoint/api/todo/queries"
	"todo-application/model"

	"github.com/gorilla/mux"
)

func Update(writer http.ResponseWriter, request *http.Request) {
	var con, _ = database.Connection()
	params := mux.Vars(request)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Invalid ID")
	}

	var todo model.Todos
	json.NewDecoder(request.Body).Decode(&todo)
	queries.UpdateTodoData(con, id, todo)
}