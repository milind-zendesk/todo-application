package queries

import (
	"database/sql"
	"todo-application/model"
)

//go:generate mockgen --build_flags=--mod=mod -package queries -destination queries_mock.go . Queries

type Queries interface {
	GetAllTodosData() ([]model.Todos, error)
	GetTodoData(id int) (model.Todos, error)
	UpdateTodoData(id int, data model.Todos) error
	InsertTodoData(data model.Todos) error
	DeleteTodoData(id int) error
}

type DBQueries struct {
	Con *sql.DB
}

func (d *DBQueries) GetAllTodosData() ([]model.Todos, error) {
	todos := []model.Todos{}

	results, err := d.Con.Query("SELECT * FROM todo;")
	if err != nil {
		return todos, err
	}

	defer results.Close()

	for results.Next() {
		var nextTodo model.Todos
		err = results.Scan(&nextTodo.Id, &nextTodo.Title, &nextTodo.Status, &nextTodo.Priority, &nextTodo.UserID)
		if err != nil {
			return todos, err
		}

		todos = append(todos, nextTodo)
	}

	return todos, nil
}

func (d *DBQueries) GetTodoData(id int) (model.Todos, error) {
	todos := model.Todos{}

	results, err := d.Con.Query("SELECT * FROM todo where id = ?", id)
	if err != nil {
		return todos, err
	}

	defer results.Close()

	for results.Next() {
		err = results.Scan(&todos.Id, &todos.Title, &todos.Status, &todos.Priority, &todos.UserID)
		if err != nil {
			return todos, err
		}
	}

	return todos, nil
}

func (d *DBQueries) UpdateTodoData(id int, data model.Todos) error {
	_, err := d.Con.Exec("UPDATE todo set title=?, status=?, priority=? where id=?", data.Title, data.Status, data.Priority, data.Id)
	if err != nil {
		return err
	}
	return nil
}

func (d *DBQueries) InsertTodoData(data model.Todos) error {
	_, err := d.Con.Exec("INSERT INTO todo(title, status, priority, user_id) VALUES (?,?,?,?)", data.Title, data.Status, data.Priority, data.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (d *DBQueries) DeleteTodoData(id int) error {
	_, err := d.Con.Exec("DELETE FROM todo where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
