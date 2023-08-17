package userqueries

import (
	"database/sql"
	"fmt"
	"todo-application/model"
)

func GetAllUsersData(con *sql.DB) ([]model.User, error) {
	users := []model.User{}

	results, err := con.Query("SELECT * FROM user;")
	if err != nil {
		return users, err
	}

	defer results.Close()

	for results.Next() {
		var nextUser model.User
		err = results.Scan(&nextUser.Id, &nextUser.Name, &nextUser.Location)
		if err != nil {
			return users, err
		}

		users = append(users, nextUser)
	}

	return users, nil
}

func GetUserData(con *sql.DB, id int) (model.User, error) {
	user := model.User{}

	results, err := con.Query("SELECT * FROM user where id = ?", id)
	if err != nil {
		return user, err
	}

	defer results.Close()

	for results.Next() {
		err = results.Scan(&user.Id, &user.Name, &user.Location)
		if err != nil {
			return user, err
		}
	}

	return user, nil
}

func GetUserTodosData(con *sql.DB, id int) (model.UserTodoDetails, error) {
	user_todo_details := model.UserTodoDetails{}

	results, err := con.Query("SELECT * FROM user where id = ?", id)
	if err != nil {
		return user_todo_details, err
	}

	defer results.Close()

	for results.Next() {
		err = results.Scan(&user_todo_details.Id, &user_todo_details.Name, &user_todo_details.Location)
		if err != nil {
			return user_todo_details, err
		}
	}

	results, err = con.Query("SELECT * FROM todo where user_id = ?", id)
	if err != nil {
		return user_todo_details, err
	}

	defer results.Close()

	priority_list := map[string]int{
		"high":   0,
		"medium": 0,
		"low":    0,
	}
	total_count := 0

	for results.Next() {
		var nextTodo model.Todos
		err = results.Scan(&nextTodo.Id, &nextTodo.Title, &nextTodo.Status, &nextTodo.Priority, &nextTodo.UserID)
		fmt.Println(nextTodo.Priority)
		if err != nil {
			return user_todo_details, err
		}

		priority_list[nextTodo.Priority] += 1
		total_count += 1
		user_todo_details.Todos = append(user_todo_details.Todos, nextTodo)
	}
	user_todo_details.TotalCount = total_count
	user_todo_details.Priorities = priority_list
	return user_todo_details, nil
}

func UpdateUserData(con *sql.DB, id int, data model.User) error {
	_, err := con.Exec("UPDATE user set name=? , location=? where id=?", data.Name, data.Location, data.Id)
	if err != nil {
		return err
	}
	return nil
}

func InsertUserData(con *sql.DB, data model.User) error {
	_, err := con.Exec("INSERT INTO user(name, location) VALUES (?,?)", data.Name, data.Location)
	if err != nil {
		return err
	}
	return nil
}
