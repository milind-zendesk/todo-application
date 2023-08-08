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
	RunServer()
}

func RunServer() {
	router := mux.NewRouter()
	router.HandleFunc("/todos", GetAllTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", GetTodo).Methods("GET")
	router.HandleFunc("/delete_todo/{id}", DeleteTodo).Methods("DELETE")
	router.HandleFunc("/edit_todo/{id}", UpdateTodo).Methods("PUT")
	router.HandleFunc("/insert_todo", CreateTodo).Methods("POST")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("There's an error with the server ", err)
	}

}

func GetAllTodos(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	todos, err := GetAllTodosData(con)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(writer).Encode(&todos)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}
}

func GetAllTodosData(con *sql.DB) ([]Todos, error) {
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

func GetTodo(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Invalid ID")
	}

	todo, err := GetTodoData(con, id)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(writer).Encode(&todo)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}
}

func GetTodoData(con *sql.DB, id int) (Todos, error) {
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
	defer result.Close()
}

func UpdateTodo(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Invalid ID")
	}

	var todo Todos
	json.NewDecoder(request.Body).Decode(&todo)
	UpdateTodoData(con, id, todo)
}

func UpdateTodoData(con *sql.DB, id int, data Todos) {
	update, err := con.Query("UPDATE todos_table set title=? , status=? where id=?", data.Title, data.Status, data.Id)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer update.Close()
	}
}

func CreateTodo(writer http.ResponseWriter, request *http.Request) {
	var todo Todos
	json.NewDecoder(request.Body).Decode(&todo)
	InsertTodoData(con, todo)
}

func InsertTodoData(con *sql.DB, data Todos) {
	insert, err := con.Query("INSERT INTO todos_table VALUES (?,?,?)", data.Id, data.Title, data.Status)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer insert.Close()
	}
}
