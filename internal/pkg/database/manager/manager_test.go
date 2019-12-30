package manager_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"

	"github.com/zinefer/habits/internal/pkg/database/manager"
)

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) TestDatabaseManager() {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")

	mock.ExpectExec("CREATE DATABASE SuperTestDatabase")
	mock.ExpectExec("DROP DATABASE SuperTestDatabase")
	// We need to setup an expectation for the contents of testfile.sql
	mock.ExpectExec(`CREATE TABLE space_vehicles [^;]+;\s+INSERT INTO space_vehicles`)

	man := manager.New(db)
	man.Create("SuperTestDatabase")
	man.Drop("SuperTestDatabase")
	man.Load("testfile.sql")

	if err := mock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}

	// Test that a bad file throws an error
	err := man.Load("nonexistantfile.sql")
	assert.NotNil(suite.T(), err)
}

func TestDatabaseManagerTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
