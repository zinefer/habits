package config

import (
	"flag"
)

// Configuration - Application config
type Configuration struct {
	ListenAddress        string
	SessionSecret        string
	GithubClientID       string
	GithubClientSecret   string
	GoogleClientID       string
	GoogleClientSecret   string
	FacebookClientID     string
	FacebookClientSecret string
}

// New - Construct a new application config
func New() *Configuration {
	c := Configuration{}
	flag.StringVar(&c.ListenAddress, "listen-addr", ":3000", "server listen address")
	flag.StringVar(&c.SessionSecret, "secret", "", "Session secret")

	flag.StringVar(&c.GithubClientID, "auth-git-id", "", "github oauth client id")
	flag.StringVar(&c.GithubClientSecret, "auth-git-secret", "", "github oauth client secret")
	flag.Parse()

	return &c
}
