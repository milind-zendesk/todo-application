package todo

import (
	"encoding/json"
	"log"
	"net/http"
	"todo-application/database"
	"todo-application/endpoint/api/todo/queries"
	"todo-application/model"
)

func Create(writer http.ResponseWriter, request *http.Request) {
	var con, _ = database.Connection()
	var todo model.Todos
	json.NewDecoder(request.Body).Decode(&todo)
	err := queries.InsertTodoData(con, todo)
	if err != nil {
		log.Fatal(err.Error())
	}
}
