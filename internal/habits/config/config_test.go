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

	var configuration *config.Configuration
	assert.NotPanics(suite.T(), func() {
		configuration = config.New()
	}, "Does not panic with a good config")

	assert.Equal(suite.T(), "127.0.0.1", configuration.Database.Host)
	assert.Equal(suite.T(), "habits_development", configuration.Database.Name)
	assert.Equal(suite.T(), "postgres", configuration.Database.Username)
}

func TestConfigurationTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
