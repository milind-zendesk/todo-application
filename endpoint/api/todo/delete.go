package todo

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"todo-application/database"
	"todo-application/endpoint/api/todo/queries"

	"github.com/gorilla/mux"
)

func Delete(writer http.ResponseWriter, request *http.Request) {
	var con, _ = database.Connection()
	params := mux.Vars(request)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Invalid ID")
	}

	err = queries.DeleteTodoData(con, id)
	if err != nil {
		log.Fatal(err.Error())
	}
}
