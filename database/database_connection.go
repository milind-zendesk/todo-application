package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {
	db, err := sql.Open("mysql", "root:TempPassword@(localhost:11335)/todo_app_db")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	} else {
		return db
	}
}
