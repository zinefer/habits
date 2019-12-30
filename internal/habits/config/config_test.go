package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/zinefer/habits/internal/habits/config"
)

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) TestConfiguration() {
	config.DatabaseConfigPath = "testconfig.yml"
	configuration := config.New()
	assert.Equal(suite.T(), "127.0.0.1", configuration.Database.Host)
	assert.Equal(suite.T(), "habits_development", configuration.Database.Name)
	assert.Equal(suite.T(), "postgres", configuration.Database.Username)
}

func (suite *TestSuite) TestDatabaseConfig() {
	config := config.NewDatabaseConfiguration()
	config.Host = "localhost"
	config.Port = "5243"
	config.Name = "moonbase"
	config.Username = "billgates"
	config.Password = "god"

	expected := "postgres://billgates:god@localhost:5243/moonbase"
	assert.Equal(suite.T(), expected, config.URI())

	config.Password = ""
	expected = "postgres://billgates@localhost:5243/moonbase"
	assert.Equal(suite.T(), expected, config.URI())

	config.Username = ""
	expected = "postgres://localhost:5243/moonbase"
	assert.Equal(suite.T(), expected, config.URI())

	config.Port = ""
	expected = "postgres://localhost/moonbase"
	assert.Equal(suite.T(), expected, config.URI())
}

func TestConfigurationTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
