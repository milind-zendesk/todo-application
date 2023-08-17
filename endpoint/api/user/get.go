package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"todo-application/database"
	userqueries "todo-application/endpoint/api/user/userQueries"

	"github.com/gorilla/mux"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	var con, _ = database.Connection()
	writer.Header().Set("Content-Type", "application/json")

	users, err := userqueries.GetAllUsersData(con)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(writer).Encode(&users)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}
}

func Get(writer http.ResponseWriter, request *http.Request) {
	var con, _ = database.Connection()
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("Invalid ID")
	}

	user, err := userqueries.GetUserData(con, id)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(writer).Encode(&user)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}
}

func GetUserTodos(writer http.ResponseWriter, request *http.Request) {
	var con, _ = database.Connection()
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("Invalid ID")
	}

	user_todos, err := userqueries.GetUserTodosData(con, id)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(writer).Encode(&user_todos)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}
}
