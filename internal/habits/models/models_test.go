package user_test

import (
	"flag"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/zinefer/habits/internal/pkg/database/manager"

	"github.com/zinefer/habits/internal/habits/helpers/test"
	"github.com/zinefer/habits/internal/habits/middlewares/database"
	
	"github.com/zinefer/habits/internal/habits/models/user"
)

type TestSuite struct {
	suite.Suite
}

var (
	postgresURI = *flag.String("psql-uri", "postgres://postgres@127.0.0.1", "Test postgres URI")
	testDB      = "habits_test"
)

var (
	conn *sqlx.DB
	dbManager *manager.DatabaseManager
	r  *chi.Mux
	db *sqlx.DB
)

func (suite *TestSuite) SetupSuite() {
	var err error
	conn, err = sqlx.Open("postgres", postgresURI)
	if err != nil {
		panic(err)
	}

	dbManager = manager.New(conn)
	dbManager.Create(testDB)

	db, err = sqlx.Open("postgres", postgresURI+"/"+testDB)
	if err != nil {
		panic(err)
	}

	man := manager.New(db)
	man.Load("schemalinkfortest.sql")
}

func (suite *TestSuite) TearDownSuite() {
	db.Close()
	dbManager.Drop(testDB)
	conn.Close()
}

func (suite *TestSuite) SetupTest() {
	r = chi.NewRouter()

	r.Use(database.DbContextMiddleware(db))
}

func (suite *TestSuite) TestUser() {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		u1 := user.New("1", "test", "admin", "jerry", "jerry@jake.com")
		_, err := u1.Save(r.Context())
		assert.NoError(suite.T(), err, "No error saving admin")

		u2 := user.New("2", "test", "admin", "jake", "jake@jerry.com")
		available, err := user.IsNameAvailable(r.Context(), u2.Name)
		assert.NoError(suite.T(), err)
		assert.False(suite.T(), available, "Unavailable name")

		u3 := user.New("2", "test", "timmy", "kim", "tim@tommy.com")
		available, err = user.IsNameAvailable(r.Context(), u3.Name)
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), available, "Unavailable name")
	})

	ts := httptest.NewServer(r)
	defer ts.Close()
	test.Request(suite.T(), ts, "GET", "/", nil)
}

func TestModelsSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
