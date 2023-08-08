package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:TempPassword@(localhost:11335)/todo_app_db")
	if err != nil {
		return nil, err
	}
	return db, nil

}
