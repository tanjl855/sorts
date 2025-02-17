package test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestShouldUpdateStats(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("update products").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("insert into product_viewers").
		WithArgs(2, 3).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if err = recordStats(db, 2, 3); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
