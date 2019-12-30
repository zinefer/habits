package manager

import (
	"io/ioutil"

	"github.com/jmoiron/sqlx"
)

// DatabaseManager manages databases
type DatabaseManager struct {
	db *sqlx.DB
}

// New DatabaseManager
func New(db *sqlx.DB) *DatabaseManager {
	return &DatabaseManager{
		db: db,
	}
}

// Create a database
func (m *DatabaseManager) Create(name string) error {
	_, err := m.db.Exec("CREATE DATABASE " + name)
	return err
}

// Drop a database
func (m *DatabaseManager) Drop(name string) error {
	_, err := m.db.Exec("DROP DATABASE " + name)
	return err
}

// Load a file into the database
func (m *DatabaseManager) Load(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	_, err = m.db.Exec(string(data))
	return err
}
