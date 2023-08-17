package endpoint

import (
	"todo-application/database"
	bulkload "todo-application/endpoint/api/bulkLoad"
	"todo-application/endpoint/api/todo"
	"todo-application/endpoint/api/todo/queries"
	"todo-application/endpoint/api/user"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()

	var con, _ = database.Connection()
	querier := &queries.DBQueries{
		Con: con,
	}

	//Bulk load random data in User and Todo Table
	router.HandleFunc("/bulk_load", bulkload.StoreFakeData).Methods("POST")

	//Routes for User
	router.HandleFunc("/users", user.GetAll).Methods("GET")
	router.HandleFunc("/users/{id}", user.Get).Methods("GET")
	router.HandleFunc("/user_todos/{id}", user.GetUserTodos).Methods("GET")
	router.HandleFunc("/edit_user/{id}", user.Update).Methods("PUT")
	router.HandleFunc("/add_users", user.Create).Methods("POST")

	//Routes for Todo
	router.HandleFunc("/todos", todo.GetAllHandlers(querier)).Methods("GET")
	router.HandleFunc("/todos/{id}", todo.GetHandlers(querier)).Methods("GET")
	router.HandleFunc("/delete_todo/{id}", todo.DeleteHandler(querier)).Methods("DELETE")
	router.HandleFunc("/edit_todo/{id}", todo.UpdateHandler(querier)).Methods("PUT")
	router.HandleFunc("/insert_todo", todo.CreateHandler(querier)).Methods("POST")

	return router
}
