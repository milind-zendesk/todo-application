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
	dbQuerier := &DBQueries{
		db,
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM todo where id = ?")).WithArgs(4)

	dbQuerier.DeleteTodoData(4)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func Test_InsertQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a database connection", err)
	}
	dbQuerier := &DBQueries{
		db,
	}
	defer db.Close()

	todo := model.Todos{
		Title:    "Work",
		Status:   "done",
		Priority: "low",
		UserID:   4,
	}

	mock.ExpectExec("INSERT INTO todo").WithArgs(todo.Title, todo.Status, todo.Priority, todo.UserID).WillReturnResult(sqlmock.NewResult(1, 1))

	dbQuerier.InsertTodoData(todo)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_UpdateQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a database connection", err)
	}
	dbQuerier := &DBQueries{
		db,
	}
	defer db.Close()

	todo := model.Todos{
		Id:       1,
		Title:    "Work",
		Status:   "done",
		Priority: "high",
	}

	mock.ExpectExec(regexp.QuoteMeta("UPDATE todo set title=?, status=?, priority=? where id=?")).WithArgs("Work", "done", "high", 1).WillReturnResult(sqlmock.NewResult(1, 1))

	dbQuerier.UpdateTodoData(1, todo)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_GetTodoData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	dbQuerier := &DBQueries{
		db,
	}
	defer db.Close()

	expected_todo := model.Todos{
		Id:       1,
		Title:    "Work",
		Status:   "done",
		Priority: "high",
		UserID:   4,
	}

	mock.
		ExpectQuery(regexp.QuoteMeta("SELECT * FROM todo where id = ?")).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "status", "priority", "user_id"}).
				AddRow(1, "Work", "done", "high", 4))

	actual_todo, err := dbQuerier.GetTodoData(1)
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
	dbQuerier := &DBQueries{
		db,
	}
	defer db.Close()

	expected_todo := []model.Todos{
		{
			Id:       1,
			Title:    "Work",
			Status:   "done",
			Priority: "low",
			UserID:   3,
		},
		{
			Id:       2,
			Title:    "Breakfast",
			Status:   "pending",
			Priority: "low",
			UserID:   2,
		},
		{
			Id:       3,
			Title:    "Exercise",
			Status:   "ongoing",
			Priority: "low",
			UserID:   1,
		},
	}

	mock.
		ExpectQuery(regexp.QuoteMeta("SELECT * FROM todo")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "status", "priority", "user_id"}).
				AddRow(1, "Work", "done", "low", 3).AddRow(2, "Breakfast", "pending", "low", 2).AddRow(3, "Exercise", "ongoing", "low", 1))

	actual_todo, err := dbQuerier.GetAllTodosData()
	if err != nil {
		t.Errorf("Something went wrong: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, expected_todo, actual_todo)
}
