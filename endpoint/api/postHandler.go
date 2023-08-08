package api

import (
	"encoding/json"
	"net/http"
	"todo-application/endpoint/api/queries"
	"todo-application/model"
)

func CreateTodo(writer http.ResponseWriter, request *http.Request) {
	var todo model.Todos
	json.NewDecoder(request.Body).Decode(&todo)
	queries.InsertTodoData(todo)
}
