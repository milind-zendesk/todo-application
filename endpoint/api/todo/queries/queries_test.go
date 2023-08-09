package queries

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_DeleteQuery(t *testing.T) {

	// 1. Prepare for the test

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	fmt.Println(db)
	fmt.Println(mock)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("DELETE from todo").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// 2. Run / Execute the test
	DeleteTodoData(db, 4)

	// 3. Check results

}
