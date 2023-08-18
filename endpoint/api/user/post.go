package user

import (
	"encoding/json"
	"log"
	"net/http"
	userqueries "todo-application/endpoint/api/user/userQueries"
	"todo-application/model"
)

func CreateHandler(queries userqueries.UserQueries) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var user model.User
		json.NewDecoder(request.Body).Decode(&user)
		err := queries.InsertUserData(user)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
