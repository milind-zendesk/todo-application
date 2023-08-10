package todo

import (
	"encoding/json"
	"net/http"
	"todo-application/database"
	"todo-application/endpoint/api/todo/queries"
	"todo-application/model"
)

func Create(writer http.ResponseWriter, request *http.Request) {
	var con, _ = database.Connection()
	var todo model.Todos
	json.NewDecoder(request.Body).Decode(&todo)
	queries.InsertTodoData(con, todo)
}
