package bulkload

import (
	"log"
	"math/rand"
	"net/http"
	"todo-application/database"
	bulkloadqueries "todo-application/endpoint/api/bulkLoad/bulkLoadQueries"
	"todo-application/model"

	"github.com/brianvoe/gofakeit/v6"
)

func StoreFakeData(writer http.ResponseWriter, request *http.Request) {
	con, err := database.Connection()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Generate 100 random user records
	users := make([]model.User, 100)
	for i := range users {
		users[i] = model.User{
			Name:     gofakeit.Name(),
			Location: gofakeit.City(),
		}
	}
	err = bulkloadqueries.BulkLoadUsers(con, users)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Generate 200 random todo tasks
	todos := make([]model.Todos, 200)
	statuses := []string{"not started", "in progress", "done"}
	priorities := []string{"low", "medium", "high"}

	for i := range todos {
		todos[i] = model.Todos{
			Title:    gofakeit.Sentence(5),
			Status:   statuses[rand.Intn(len(statuses))],
			Priority: priorities[rand.Intn(len(priorities))],
			UserID:   rand.Intn(100) + 1,
		}
	}
	err = bulkloadqueries.BulkLoadTodos(con, todos)
	if err != nil {
		log.Fatal(err.Error())
	}
}
