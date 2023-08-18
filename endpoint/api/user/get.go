package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	userqueries "todo-application/endpoint/api/user/userQueries"

	"github.com/gorilla/mux"
)

func GetAllHandler(queries userqueries.UserQueries) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		users, err := queries.GetAllUsersData()
		if err != nil {
			panic(err)
		}

		err = json.NewEncoder(writer).Encode(&users)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
	}
}

func GetHandler(queries userqueries.UserQueries) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		params := mux.Vars(request)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatal("Invalid ID")
		}

		user, err := queries.GetUserData(id)
		if err != nil {
			panic(err)
		}

		err = json.NewEncoder(writer).Encode(&user)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
	}
}

func GetUserTodosHandler(queries userqueries.UserQueries) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		params := mux.Vars(request)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatal("Invalid ID")
		}

		user_todos, err := queries.GetUserTodosData(id)
		if err != nil {
			panic(err)
		}

		err = json.NewEncoder(writer).Encode(&user_todos)
		if err != nil {
			log.Fatalln("There was an error encoding the initialized struct")
		}
	}
}
