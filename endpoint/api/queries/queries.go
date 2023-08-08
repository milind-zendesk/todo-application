package queries

import (
	"fmt"
	"todo-application/database"
	"todo-application/model"
)

var con, _ = database.Connection()

func GetAllTodosData() ([]model.Todos, error) {
	todos := []model.Todos{}

	results, err := con.Query("SELECT * FROM todos_table;")
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

	results, err := con.Query("SELECT * FROM todos_table where id = ?", id)
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
	update, err := con.Query("UPDATE todos_table set title=? , status=? where id=?", data.Title, data.Status, data.Id)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer update.Close()
	}
}

func InsertTodoData(data model.Todos) {
	insert, err := con.Query("INSERT INTO todos_table VALUES (?,?,?)", data.Id, data.Title, data.Status)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer insert.Close()
	}
}

func DeleteTodoData(id int) {
	result, err := con.Query("DELETE FROM todos_table where id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
	defer result.Close()
}
