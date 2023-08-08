package endpoint

import (
	"log"
	"net/http"
	"todo-application/endpoint/api"

	"github.com/gorilla/mux"
)

func RunServer() {
	router := mux.NewRouter()
	router.HandleFunc("/todos", api.GetAllTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", api.GetTodo).Methods("GET")
	router.HandleFunc("/delete_todo/{id}", api.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/edit_todo/{id}", api.UpdateTodo).Methods("PUT")
	router.HandleFunc("/insert_todo", api.CreateTodo).Methods("POST")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("There's an error with the server ", err)
	}

}
