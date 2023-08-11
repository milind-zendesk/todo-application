package bulkloadqueries

import (
	"database/sql"
	"todo-application/model"
)

func BulkLoadUsers(con *sql.DB, users []model.User) error {
	for _, user := range users {
		insert, err := con.Query("INSERT INTO user(name, location) VALUES (?,?)", user.Name, user.Location)
		if err != nil {
			return err
		} else {
			insert.Close()
		}
	}
	return nil
}

func BulkLoadTodos(con *sql.DB, todos []model.Todos) error {
	for _, todo := range todos {
		insert, err := con.Query("INSERT INTO todo(title, status, priority, user_id) VALUES (?,?,?,?)", todo.Title, todo.Status, todo.Priority, todo.UserID)
		if err != nil {
			return err
		} else {
			insert.Close()
		}
	}
	return nil
}
