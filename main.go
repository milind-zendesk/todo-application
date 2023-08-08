package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"todo-application/database"

	"github.com/gorilla/mux"
)

var con = database.Connection()

type Todos struct {
	Id     int
	Title  string
	Status string
}

func main() {
	fmt.Println("Hello World!")
	RunServer()

	// insert, err := con.Query("INSERT INTO todos_table VALUES ( 2, 'Deployment', 'Running' )")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()
	// defer con.Close()
}

func RunServer() {
	router := mux.NewRouter()
	router.HandleFunc("/todos", getAllTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", getTodo).Methods("GET")
	router.HandleFunc("/delete_todo/{id}", DeleteTodo).Methods("DELETE")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("There's an error with the server ", err)
	}

}

func getAllTodos(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	todos, err := getAllTodosData(con)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(writer).Encode(&todos)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}
}

func getAllTodosData(con *sql.DB) ([]Todos, error) {
	todos := []Todos{}

	results, err := con.Query("SELECT * FROM todos_table;")
	if err != nil {
		return todos, err
	}

	defer results.Close()

	for results.Next() {
		var nextTodo Todos
		err = results.Scan(&nextTodo.Id, &nextTodo.Title, &nextTodo.Status)
		if err != nil {
			return todos, err
		}

		todos = append(todos, nextTodo)
	}

	return todos, nil
}

func getTodo(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Invalid ID")
	}

	todo, err := getTodoData(con, id)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(writer).Encode(&todo)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}
}

func getTodoData(con *sql.DB, id int) (Todos, error) {
	todos := Todos{}

	results, err := con.Query("SELECT * FROM todos_table where id = ?", id)
	if err != nil {
		return todos, err
	}

	defer results.Close()

	for results.Next() {
		err = results.Scan(&todos.Id, &todos.Title, &todos.Status)
		if err != nil {
			return todos, err
		}
	}

	return todos, nil
}

func DeleteTodo(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Invalid ID")
	}

	result, err := con.Query("DELETE FROM todos_table where id = ?", id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("done")
	defer result.Close()
}
