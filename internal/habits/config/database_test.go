package config_test

import (
	"github.com/stretchr/testify/assert"

	"github.com/zinefer/habits/internal/habits/config"
)

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
