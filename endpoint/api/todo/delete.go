package todo

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"todo-application/endpoint/api/todo/queries"

	"github.com/gorilla/mux"
)

func DeleteHandler(queries queries.Queries) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)

		id, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Println("Invalid ID")
		}

		err = queries.DeleteTodoData(id)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
