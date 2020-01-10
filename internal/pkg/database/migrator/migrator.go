package migrator

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/zinefer/habits/internal/pkg/database/manager"
)

// SQLMigrator migrates a database against a collection of sql migrations
type SQLMigrator struct {
	migrationsPath string
	db             *sqlx.DB
}

const schema string = `CREATE TABLE schema_migrations (
    version VARCHAR(14) UNIQUE NOT NULL, /* YYYYMMDDHHMMSS */
);`

// New returns a SqlMigrator
func New(db *sqlx.DB, migrationsPath string) *SQLMigrator {
	return &SQLMigrator{
		migrationsPath: migrationsPath,
		db:             db,
	}
}

// Migrate reconciles a database to a state described by a collection of sql migrations
func (m *SQLMigrator) Migrate() error {
	if !m.MigrationsTableExists() {
		m.initializeDatabase()
	}

	files, err := ioutil.ReadDir(m.migrationsPath)
	if err != nil {
		return err
	}

	tx, err := m.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	dbManager := manager.New(tx)

	for _, f := range files {
		name := f.Name()
		version := getVersionFromFilename(name)
		needsRun := m.hasMigrationAlreadyRun(version) == false

		if needsRun {
			fmt.Printf("Executing new migration %v▲\n", name)

			path := filepath.Join(m.migrationsPath, name, "up.sql")
			err = dbManager.Load(path)
			if err != nil {
				tx.Rollback()
				return err
			}

			err = markMigrationAlreadyRun(tx, version)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit()
}

// Rollback to the last version
func (m *SQLMigrator) Rollback() error {
	if !m.MigrationsTableExists() {
		return errors.New("Migrations table does not exist")
	}

	files, err := ioutil.ReadDir(m.migrationsPath)
	if err != nil {
		return err
	}

	last := m.getLastRunMigration()
	for h := range files {
		i := len(files) - 1 - h
		file := files[i]
		version := getVersionFromFilename(file.Name())

		if version == last {
			tx, err := m.db.Begin()
			if err != nil {
				tx.Rollback()
				return err
			}

			fmt.Printf("Rolling back last run migration %v▼\n", file.Name())

			dbManager := manager.New(tx)
			path := filepath.Join(m.migrationsPath, file.Name(), "down.sql")
			err = dbManager.Load(path)
			if err != nil {
				tx.Rollback()
				return err
			}

			err = markMigrationNotRun(tx, version)
			if err != nil {
				tx.Rollback()
				return err
			}

			return tx.Commit()
		}
	}

	return err
}

// StubNewMigration creates a migration stub
func (m *SQLMigrator) StubNewMigration(name string) (string, error) {
	version := time.Now().Format("20060102150405")
	folder := version + "_" + name
	path := filepath.Join(m.migrationsPath, folder)

	exists, err := exists(path)
	if err != nil {
		return version, err
	}

	if exists {
		return version, errors.New("Migration already exists")
	}

	err = os.Mkdir(path, 0744)
	if err != nil {
		return version, err
	}

	_, err = os.Create(filepath.Join(path, "up.sql"))
	if err != nil {
		return version, err
	}

	_, err = os.Create(filepath.Join(path, "down.sql"))
	if err != nil {
		return version, err
	}

	return version, nil
}

func (m *SQLMigrator) initializeDatabase() (bool, error) {
	m.db.MustExec(schema)
	return true, nil
}

// MigrationsTableExists returns true if a migrations table exists
func (m *SQLMigrator) MigrationsTableExists() bool {
	_, err := m.db.Exec("SELECT 1 FROM schema_migrations LIMIT 1")
	return err == nil
}

func (m *SQLMigrator) hasMigrationAlreadyRun(version string) bool {
	var result bool
	q := m.db.QueryRow("SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE version = $1)", version)
	err := q.Scan(&result)
	if err != nil {
		panic(err)
	}
	return result
}

func (m *SQLMigrator) getLastRunMigration() string {
	var result string
	q := m.db.QueryRow("SELECT version FROM schema_migrations ORDER BY version DESC LIMIT 1")
	err := q.Scan(&result)
	if err != nil {
		panic(err)
	}
	return result
}

func markMigrationNotRun(db sqlx.Execer, version string) error {
	_, err := db.Exec("DELETE FROM schema_migrations WHERE version = $1 LIMIT 1", version)
	return err
}

func markMigrationAlreadyRun(db sqlx.Execer, version string) error {
	_, err := db.Exec("INSERT INTO schema_migrations (version) VALUES ($1)", version)
	return err
}

func getVersionFromFilename(filename string) string {
	split := strings.Split(filename, "_")
	return split[0]
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
