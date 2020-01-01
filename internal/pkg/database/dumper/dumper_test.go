package dumper_test

import (
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/zinefer/habits/internal/pkg/database/dumper"
	"github.com/zinefer/habits/internal/pkg/database/manager"
)

type TestSuite struct {
	suite.Suite
}

var (
	postgresURI = *flag.String("psql-uri", "postgres://postgres@127.0.0.1", "Test postgres URI")
	testData    = "testdata.sql"
	dumpFile    = "test.dump"
	testDB      = "dumper_test"
)

func (suite *TestSuite) TestSQLDumper() {
	conn, err := sqlx.Open("postgres", postgresURI)
	if err != nil {
		panic(err)
	}

	dbManager := manager.New(conn)
	err = dbManager.Create(testDB)
	defer conn.Close()
	assert.NoError(suite.T(), err, "Created with no error")

	db, err := sqlx.Open("postgres", postgresURI+"/"+testDB)
	if err != nil {
		panic(err)
	}

	testManager := manager.New(db)
	err = testManager.Load(testData)
	assert.NoError(suite.T(), err, "Loaded with no error")

	dump := dumper.New(db)
	err = dump.Dump(dumpFile)
	assert.NoError(suite.T(), err, "Dumped with no error")

	expected, _ := ioutil.ReadFile(testData)
	dumped, _ := ioutil.ReadFile(dumpFile)

	assert.Equal(suite.T(), string(expected), string(dumped))

	db.Close()

	err = dbManager.Drop(testDB)
	assert.NoError(suite.T(), err, "Dropped with no error")

	err = os.Remove(dumpFile)
	assert.NoError(suite.T(), err, "Deleted with no error")
}

func TestSQLDumperTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
