package api

import (
	"fmt"
	"net/http"
	"strconv"
	"todo-application/endpoint/api/queries"

	"github.com/gorilla/mux"
)

func DeleteTodo(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Invalid ID")
	}

	queries.DeleteTodoData(id)
}
