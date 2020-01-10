package config

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	// CertsConfigPath points to a directory we will cache LTS certs in
	CertsConfigPath = "certs"
	// SecretConfigPath points to the session secret file
	SecretConfigPath = "secret"
	// DatabaseConfigPath points to the database config yaml
	DatabaseConfigPath = "database/config.yml"
	// DatabaseSchemaPath points to the database schema sql
	DatabaseSchemaPath = "database/schema.sql"
	// DatabaseMigrationPath points to the database migrations
	DatabaseMigrationPath = "database/migrations"
)

// Configuration - Application config
type Configuration struct {
	Hostname      string
	Environment   string
	ListenAddress string
	SessionSecret []byte
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

	env := os.Getenv("HABITS_ENVIRONMENT")
	if len(env) == 0 {
		env = "development"
	}

	flag.StringVar(&c.Hostname, "hostname", "habits.watch", "Application hostname")
	flag.StringVar(&c.Environment, "env", env, "Environment to run application in")
	flag.StringVar(&c.ListenAddress, "listen-addr", "80", "server listen address")

	flag.StringVar(&c.GithubClientID, "auth-github-id", os.Getenv("HABITS_OAUTH_GITHUB_ID"), "github oauth client id")
	flag.StringVar(&c.GithubClientSecret, "auth-github-secret", os.Getenv("HABITS_OAUTH_GITHUB_SECRET"), "github oauth client secret")
	flag.StringVar(&c.GoogleClientID, "auth-google-id", os.Getenv("HABITS_OAUTH_GOOGLE_ID"), "google oauth client id")
	flag.StringVar(&c.GoogleClientSecret, "auth-google-secret", os.Getenv("HABITS_OAUTH_GOOGLE_SECRET"), "google oauth client secret")
	flag.StringVar(&c.FacebookClientID, "auth-facebook-id", os.Getenv("HABITS_OAUTH_FACEBOOK_ID"), "facebook oauth client id")
	flag.StringVar(&c.FacebookClientSecret, "auth-facebook-secret", os.Getenv("HABITS_OAUTH_FACEBOOK_SECRET"), "facebook oauth client secret")

	flag.Parse()

	c.parseDatabaseConfig()
	c.SessionSecret = c.readSecretConfig()

	dbHost := os.Getenv("HABITS_DATABASE_HOST")
	if len(dbHost) > 0 {
		c.Database.Host = dbHost
	}

	dbUser := os.Getenv("HABITS_DATABASE_USER")
	if len(dbUser) > 0 {
		c.Database.Username = dbUser
	}

	dbPassword := os.Getenv("HABITS_DATABASE_PASSWORD")
	if len(dbPassword) > 0 {
		c.Database.Password = dbPassword
	}

	return &c
}

func (c *Configuration) readSecretConfig() []byte {
	data, err := ioutil.ReadFile(SecretConfigPath)
	if err != nil || len(data) == 0 {
		if c.IsProduction() {
			fmt.Println("Creating secret data for encrypting session")
			data = c.CreateSecretConfig()
		}
	}
	return data
}

// CreateSecretConfig writes a random secret config
func (c *Configuration) CreateSecretConfig() []byte {
	data := generateRandomKey()
	file, err := os.Create(SecretConfigPath)
	_, err = file.Write(data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

// IsProduction returns true if the application is running in production
func (c *Configuration) IsProduction() bool {
	return c.Environment == "production"
}

func (c *Configuration) parseDatabaseConfig() {
	source, err := ioutil.ReadFile(DatabaseConfigPath)
	if err != nil {
		panic(err)
	}
	configs := make(map[string]DatabaseConfiguration)
	err = yaml.Unmarshal(source, configs)
	if err != nil {
		panic(err)
	}
	config := configs[c.Environment]
	c.Database = &config
}

func generateRandomKey() []byte {
	k := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		return nil
	}
	return k
}
