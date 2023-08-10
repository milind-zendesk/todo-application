package queries

import (
	"database/sql"
	"todo-application/model"
)

type Queries interface {
	GetAllTodosData(con *sql.DB)
	GetTodoData(con *sql.DB, id int)
	UpdateTodoData(con *sql.DB, id int, data model.Todos)
	InsertTodoData(con *sql.DB, data model.Todos)
	DeleteTodoData(con *sql.DB, id int)
}

func GetAllTodosData(con *sql.DB) ([]model.Todos, error) {
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

func GetTodoData(con *sql.DB, id int) (model.Todos, error) {
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

func UpdateTodoData(con *sql.DB, id int, data model.Todos) error {
	_, err := con.Exec("UPDATE todo set title=? , status=? where id=?", data.Title, data.Status, data.Id)
	if err != nil {
		return err
	}
	return nil
}

func InsertTodoData(con *sql.DB, data model.Todos) error {
	_, err := con.Exec("INSERT INTO todo(title, status) VALUES (?,?)", data.Title, data.Status)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTodoData(con *sql.DB, id int) error {
	_, err := con.Exec("DELETE FROM todo where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
