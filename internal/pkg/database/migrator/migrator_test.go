package migrator_test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"

	"github.com/zinefer/habits/internal/pkg/database/migrator"
)

var testMigrations = "test_migrations"

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) TestDatabaseManager() {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")

	mock.ExpectExec("SELECT 1 FROM schema_migrations LIMIT 1").
		WillReturnError(errors.New("schema_migrations does not exist"))

	mock.ExpectExec("CREATE TABLE schema_migrations").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectBegin()

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM schema_migrations WHERE version").
		WithArgs("20191230183137").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	mock.ExpectExec("CREATE TABLE users").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO schema_migrations \\(version\\) VALUES").
		WithArgs("20191230183137").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM schema_migrations WHERE version").
		WithArgs("20191231115542").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	mock.ExpectExec("CREATE TABLE account").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO schema_migrations \\(version\\) VALUES").
		WithArgs("20191231115542").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM schema_migrations WHERE version").
		WithArgs("20191231120950").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	mock.ExpectExec("CREATE INDEX idx_account_last_login").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO schema_migrations \\(version\\) VALUES").
		WithArgs("20191231120950").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	// Rollback
	mock.ExpectExec("SELECT 1 FROM schema_migrations LIMIT 1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery("SELECT version FROM schema_migrations ORDER BY version DESC LIMIT 1").
		WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("20191231120950"))

	mock.ExpectBegin()

	mock.ExpectExec("DROP INDEX idx_account_last_login").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("DELETE FROM schema_migrations WHERE version =").
		WithArgs("20191231120950").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	migrate := migrator.New(db, testMigrations)

	// Test migrating
	err := migrate.Migrate()
	assert.NoError(suite.T(), err, "Migrated with no error")

	err = migrate.Rollback()
	assert.NoError(suite.T(), err, "Rolledback with no error")

	// Test stubbing
	version, err := migrate.StubNewMigration("test")
	assert.NoError(suite.T(), err, "Stubbed with no error")

	folder := version + "_test"
	testPath := filepath.Join(testMigrations, folder)

	assert.FileExists(suite.T(), filepath.Join(testPath, "up.sql"))
	assert.FileExists(suite.T(), filepath.Join(testPath, "down.sql"))

	err = removeContents(testPath)
	assert.NoError(suite.T(), err, "Deleted test migration children with no error")
	err = os.Remove(testPath)
	assert.NoError(suite.T(), err, "Stubbed with no error")

	if err := mock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDatabaseManagerTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func removeContents(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return err
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}
	return nil
}
