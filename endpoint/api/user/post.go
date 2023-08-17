package user

import (
	"encoding/json"
	"log"
	"net/http"
	"todo-application/database"
	userqueries "todo-application/endpoint/api/user/userQueries"
	"todo-application/model"
)

func Create(writer http.ResponseWriter, request *http.Request) {
	var con, _ = database.Connection()
	var user model.User
	json.NewDecoder(request.Body).Decode(&user)
	err := userqueries.InsertUserData(con, user)
	if err != nil {
		log.Fatal(err.Error())
	}
}
