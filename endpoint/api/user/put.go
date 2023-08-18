package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	userqueries "todo-application/endpoint/api/user/userQueries"
	"todo-application/model"

	"github.com/gorilla/mux"
)

func UpdateHandler(queries userqueries.UserQueries) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)

		id, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Println("Invalid ID")
		}

		var user model.User
		json.NewDecoder(request.Body).Decode(&user)
		err = queries.UpdateUserData(id, user)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
