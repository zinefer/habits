package database_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/zinefer/habits/internal/habits/helpers/test"
	"github.com/zinefer/habits/internal/habits/middlewares/database"
)

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) TestDbContextMiddleware() {
	r := chi.NewRouter()
	odb, _ := sqlx.Open("sqlite3", ":memory:")

	r.Use(database.DbContextMiddleware(odb))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		db := database.GetDbFromContext(r.Context())
		assert.Equal(suite.T(), odb, db)
	})

	ts := httptest.NewServer(r)
	defer ts.Close()

	test.Request(suite.T(), ts, "GET", "/", nil)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
