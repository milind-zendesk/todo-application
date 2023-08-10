package queries

import (
	"database/sql"
	"fmt"
	"todo-application/database"
	"todo-application/model"
)

var con, _ = database.Connection()

func GetAllTodosData() ([]model.Todos, error) {
	todos := []model.Todos{}

	results, err := con.Query("SELECT * FROM todo;")
	if err != nil {
		return todos, err
	}

	defer results.Close()

	for results.Next() {
		var nextTodo model.Todos
		err = results.Scan(&nextTodo.Id, &nextTodo.Title, &nextTodo.Status)
		if err != nil {
			return todos, err
		}

		todos = append(todos, nextTodo)
	}

	return todos, nil
}

func GetTodoData(id int) (model.Todos, error) {
	todos := model.Todos{}

	results, err := con.Query("SELECT * FROM todo where id = ?", id)
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

func UpdateTodoData(id int, data model.Todos) {
	update, err := con.Query("UPDATE todo set title=? , status=? where id=?", data.Title, data.Status, data.Id)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer update.Close()
	}
}

func InsertTodoData(con *sql.DB, data model.Todos) {
	_, err := con.Exec("INSERT INTO todo(title, status) VALUES (?,?)", data.Title, data.Status)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func DeleteTodoData(con *sql.DB, id int) {
	_, err := con.Exec("DELETE FROM todo where id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
}
