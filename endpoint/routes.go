package endpoint

import (
	"todo-application/endpoint/api/todo"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/todos", todo.GetAll).Methods("GET")
	router.HandleFunc("/todos/{id}", todo.Get).Methods("GET")
	router.HandleFunc("/delete_todo/{id}", todo.Delete).Methods("DELETE")
	router.HandleFunc("/edit_todo/{id}", todo.Update).Methods("PUT")
	router.HandleFunc("/insert_todo", todo.Create).Methods("POST")

	return router
}
