package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"todo-application/database"
	userqueries "todo-application/endpoint/api/user/userQueries"
	"todo-application/model"

	"github.com/gorilla/mux"
)

func Update(writer http.ResponseWriter, request *http.Request) {
	var con, _ = database.Connection()
	params := mux.Vars(request)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Invalid ID")
	}

	var user model.User
	json.NewDecoder(request.Body).Decode(&user)
	err = userqueries.UpdateUserData(con, id, user)
	if err != nil {
		log.Fatal(err.Error())
	}
}
