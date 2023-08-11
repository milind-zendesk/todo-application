package userqueries

import (
	"database/sql"
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
