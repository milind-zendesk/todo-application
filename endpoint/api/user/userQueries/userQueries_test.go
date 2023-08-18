package userqueries

import (
	"regexp"
	"testing"
	"todo-application/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a database connection", err)
	}
	dbConn := &DBConn{
		db,
	}
	defer db.Close()

	expected_user := []model.User{
		{
			Id:       1,
			Name:     "Milind Shinde",
			Location: "Pune",
		},
		{
			Id:       2,
			Name:     "Devesh Chinchole",
			Location: "Jalgaon",
		},
		{
			Id:       3,
			Name:     "Luke Josh",
			Location: "Melbourn",
		},
	}

	mock.
		ExpectQuery(regexp.QuoteMeta("SELECT * FROM user")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "location"}).
				AddRow(1, "Milind Shinde", "Pune").AddRow(2, "Devesh Chinchole", "Jalgaon").AddRow(3, "Luke Josh", "Melbourn"))

	actual_user, err := dbConn.GetAllUsersData()
	if err != nil {
		t.Errorf("Something went wrong: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, expected_user, actual_user)
}

func Test_GetUserData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	dbConn := &DBConn{
		db,
	}
	defer db.Close()

	expected_user := model.User{
		Id:       1,
		Name:     "Milind Shinde",
		Location: "Pune",
	}

	mock.
		ExpectQuery(regexp.QuoteMeta("SELECT * FROM user where id = ?")).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "location"}).
				AddRow(1, "Milind Shinde", "Pune"))

	actual_user, err := dbConn.GetUserData(1)
	if err != nil {
		t.Errorf("Something went wrong: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, expected_user, actual_user)
}

func Test_UpdateQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a database connection", err)
	}
	dbConn := &DBConn{
		db,
	}
	defer db.Close()

	user := model.User{
		Id:       1,
		Name:     "Milind Shinde",
		Location: "Pune",
	}

	mock.ExpectExec(regexp.QuoteMeta("UPDATE user set name=? , location=? where id=?")).WithArgs(user.Name, user.Location, 1).WillReturnResult(sqlmock.NewResult(1, 1))

	dbConn.UpdateUserData(1, user)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_InsertQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a database connection", err)
	}
	dbQuerier := &DBConn{
		db,
	}
	defer db.Close()

	user := model.User{
		Name:     "Milind Shinde",
		Location: "Pune",
	}

	mock.ExpectExec("INSERT INTO user").WithArgs(user.Name, user.Location).WillReturnResult(sqlmock.NewResult(1, 1))

	dbQuerier.InsertUserData(user)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
