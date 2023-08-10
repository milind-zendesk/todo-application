package queries

import (
	"testing"
	"todo-application/model"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_DeleteQuery(t *testing.T) {

	// 1. Prepare for the test

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("DELETE from todo")
	mock.ExpectCommit()

	// 2. Run / Execute the test
	DeleteTodoData(db, 4)

	// 3. Check results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func Test_InsertQuery(t *testing.T) {
	//Prepare the tests
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	todo := model.Todos{
		Title:  "Work",
		Status: "Done",
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO todo").WithArgs(todo.Title, todo.Status).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	//Execute the tests
	InsertTodoData(db, todo)

	// Check the results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
