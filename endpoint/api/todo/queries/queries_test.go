package queries

import (
	"regexp"
	"testing"
	"todo-application/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_DeleteQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM todo where id = ?").WithArgs(4)

	DeleteTodoData(db, 4)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func Test_InsertQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	todo := model.Todos{
		Title:  "Work",
		Status: "Done",
	}

	mock.ExpectExec("INSERT INTO todo").WithArgs(todo.Title, todo.Status).WillReturnResult(sqlmock.NewResult(1, 1))

	InsertTodoData(db, todo)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_UpdateQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	todo := model.Todos{
		Id:     1,
		Title:  "Work",
		Status: "Done",
	}

	mock.ExpectExec(regexp.QuoteMeta("UPDATE todo set title=? , status=? where id=?")).WithArgs("Work", "Done", 1).WillReturnResult(sqlmock.NewResult(1, 1))

	UpdateTodoData(db, 1, todo)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_GetTodoData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	expected_todo := model.Todos{
		Id:     1,
		Title:  "Work",
		Status: "Done",
	}

	mock.
		ExpectQuery(regexp.QuoteMeta("SELECT * FROM todo where id = ?")).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "status"}).
				AddRow(1, "Work", "Done"))

	actual_todo, err := GetTodoData(db, 1)
	if err != nil {
		t.Errorf("Something went wrong: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, expected_todo, actual_todo)
}

func Test_GetAllTodoData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	expected_todo := []model.Todos{
		{
			Id:     1,
			Title:  "Work",
			Status: "Done",
		},
		{
			Id:     2,
			Title:  "Breakfast",
			Status: "Pending",
		},
		{
			Id:     3,
			Title:  "Exercise",
			Status: "Ongoing",
		},
	}

	mock.
		ExpectQuery(regexp.QuoteMeta("SELECT * FROM todo")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "status"}).
				AddRow(1, "Work", "Done").AddRow(2, "Breakfast", "Pending").AddRow(3, "Exercise", "Ongoing"))

	actual_todo, err := GetAllTodosData(db)
	if err != nil {
		t.Errorf("Something went wrong: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, expected_todo, actual_todo)
}
