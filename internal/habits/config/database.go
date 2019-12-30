package config

import "fmt"

// DatabaseConfiguration to configure the database
type DatabaseConfiguration struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
}

// NewDatabaseConfiguration returns a new DatabaseConfiguration
func NewDatabaseConfiguration() *DatabaseConfiguration {
	return &DatabaseConfiguration{}
}

// URI describes a postgres connection
func (db *DatabaseConfiguration) URI() string {
	user := db.Username
	if len(user) > 0 && len(db.Password) > 0 {
		user = fmt.Sprintf("%v:%v", user, db.Password)
	}

	if len(user) > 0 {
		user = user + "@"
	}

	host := db.Host
	if len(host) > 0 && len(db.Port) > 0 {
		host = fmt.Sprintf("%v:%v", host, db.Port)
	}

	return fmt.Sprintf("postgres://%v%v/%v", user, host, db.Name)
}
