package todo

import (
	"fmt"
	"net/http"
	"strconv"
	"todo-application/database"
	"todo-application/endpoint/api/todo/queries"

	"github.com/gorilla/mux"
)

var con, _ = database.Connection()

func Delete(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Invalid ID")
	}

	queries.DeleteTodoData(con, id)
}
