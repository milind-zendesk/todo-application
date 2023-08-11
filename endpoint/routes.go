package endpoint

import (
	bulkload "todo-application/endpoint/api/bulkLoad"
	"todo-application/endpoint/api/todo"
	"todo-application/endpoint/api/user"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()

	//Bulk load random data in User and Todo Table
	router.HandleFunc("/bulk_load", bulkload.StoreFakeData).Methods("POST")

	//Routes for User
	router.HandleFunc("/users", user.GetAll).Methods("GET")
	router.HandleFunc("/users/{id}", user.Get).Methods("GET")

	//Routes for Todo
	router.HandleFunc("/todos", todo.GetAll).Methods("GET")
	router.HandleFunc("/todos/{id}", todo.Get).Methods("GET")
	router.HandleFunc("/delete_todo/{id}", todo.Delete).Methods("DELETE")
	router.HandleFunc("/edit_todo/{id}", todo.Update).Methods("PUT")
	router.HandleFunc("/insert_todo", todo.Create).Methods("POST")

	return router
}
