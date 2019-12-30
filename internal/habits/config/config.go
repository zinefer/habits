package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var (
	// DatabaseConfigPath points to the database config yaml
	DatabaseConfigPath = "database/config.yml"
	// DatabaseSchemaPath points to the database schema sql
	DatabaseSchemaPath = "database/schema.sql"
)

// Configuration - Application config
type Configuration struct {
	ListenAddress string
	SessionSecret string
	// oAuth
	GithubClientID       string
	GithubClientSecret   string
	GoogleClientID       string
	GoogleClientSecret   string
	FacebookClientID     string
	FacebookClientSecret string
	// Database
	Database *DatabaseConfiguration
}

// New - Construct a new application config
func New() *Configuration {
	c := Configuration{
		Database: NewDatabaseConfiguration(),
	}

	flag.StringVar(&c.ListenAddress, "listen-addr", ":3000", "server listen address")
	flag.StringVar(&c.SessionSecret, "secret", "", "Session secret")

	flag.StringVar(&c.GithubClientID, "auth-git-id", "97bbebbbb9b0e0e89130", "github oauth client id")
	flag.StringVar(&c.GithubClientSecret, "auth-git-secret", "f3dbb822eeac3a6c3ecf2297b8b75f4dd7700df0", "github oauth client secret")

	flag.Parse()

	c.parseDatabaseConfig()

	return &c
}

func (c *Configuration) parseDatabaseConfig() {
	source, err := ioutil.ReadFile(DatabaseConfigPath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, c.Database)
	if err != nil {
		panic(err)
	}
}
