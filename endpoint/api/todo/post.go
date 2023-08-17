package todo

import (
	"encoding/json"
	"log"
	"net/http"
	"todo-application/endpoint/api/todo/queries"
	"todo-application/model"
)

func CreateHandler(queries queries.Queries) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var todo model.Todos
		json.NewDecoder(request.Body).Decode(&todo)
		err := queries.InsertTodoData(todo)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
