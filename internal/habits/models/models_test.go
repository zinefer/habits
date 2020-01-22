package user_test

import (
	"flag"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/zinefer/habits/internal/pkg/database/manager"

	"github.com/zinefer/habits/internal/habits/helpers/test"
	"github.com/zinefer/habits/internal/habits/middlewares/database"

	"github.com/zinefer/habits/internal/habits/models/activity"
	"github.com/zinefer/habits/internal/habits/models/deployment"
	"github.com/zinefer/habits/internal/habits/models/habit"
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
	conn      *sqlx.DB
	dbManager *manager.DatabaseManager
	r         *chi.Mux
	db        *sqlx.DB
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

func (suite *TestSuite) TestDeployment() {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		d1 := deployment.New("abcdefg")
		err := d1.Save(r.Context())
		assert.NoError(suite.T(), err, "Saved Deployment with no errors")
		assert.NotEqual(suite.T(), time.Time{}, d1.Created)

		d2, err := deployment.FindByVersion(r.Context(), "abcdefg")
		assert.NoError(suite.T(), err, "Found Deployment by version with no errors")
		assert.Equal(suite.T(), d1, d2)

		present, err := deployment.VersionPresent(r.Context(), "abcdefg")
		assert.NoError(suite.T(), err, "VersionPresent with no errors")
		assert.Equal(suite.T(), true, present)

		present, err = deployment.VersionPresent(r.Context(), "nope")
		assert.NoError(suite.T(), err, "VersionPresent with no errors")
		assert.Equal(suite.T(), false, present)
	})

	ts := httptest.NewServer(r)
	defer ts.Close()
	test.Request(suite.T(), ts, "GET", "/", nil)
}

func (suite *TestSuite) TestUser() {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		u1 := user.New("1", "test", "admin", "jerry", "jerry@jake.com")
		err := u1.Save(r.Context())
		assert.NoError(suite.T(), err, "No error saving admin")
		assert.NotZero(suite.T(), u1.ID)

		u2 := user.New("2", "test", "admin", "jake", "jake@jerry.com")
		available, err := user.IsNameAvailable(r.Context(), u2.Name)
		assert.NoError(suite.T(), err)
		assert.False(suite.T(), available, "Unavailable name")

		u3 := user.New("2", "test", "timmy", "kim", "tim@tommy.com")
		available, err = user.IsNameAvailable(r.Context(), u3.Name)
		assert.NoError(suite.T(), err)
		assert.True(suite.T(), available, "Unavailable name")

		admin, err := user.FindByID(r.Context(), u1.ID)
		assert.NoError(suite.T(), err)
		u1.Created = admin.Created
		assert.Equal(suite.T(), u1, admin)

		u4 := user.New("2", "test", "timmy timmer", "kim", "tim@tommy.com")
		assert.Equal(suite.T(), "timmy-timmer", u4.Name)

		u5 := user.New("2", "test", "()#HarryPotter!$!", "kim", "tim@tommy.com")
		assert.Equal(suite.T(), "harrypotter", u5.Name)
	})

	ts := httptest.NewServer(r)
	defer ts.Close()
	test.Request(suite.T(), ts, "GET", "/", nil)
}

func (suite *TestSuite) TestHabit() {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		u := user.New("2", "test", "habiter", "quitin", "quitin@fillory.com")
		u.Save(r.Context())

		h := habit.New(u.ID, "habit testing")
		err := h.Save(r.Context())
		assert.NoError(suite.T(), err, "Saved habit with no errors")
	})

	ts := httptest.NewServer(r)
	defer ts.Close()
	test.Request(suite.T(), ts, "GET", "/", nil)
}

func (suite *TestSuite) TestActivity() {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		u := user.New("3", "test", "activiter", "elliot", "elliot@fillory.com")
		u.Save(r.Context())
		h := habit.New(u.ID, "habit testing")
		h.Save(r.Context())

		a1 := activity.New(h.ID, time.Time{}, -7)
		err := a1.Save(r.Context())
		assert.NoError(suite.T(), err, "Saved Activity (zero moment) with no errors")

		moment := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
		a2 := activity.New(h.ID, moment, -7)
		err = a2.Save(r.Context())
		assert.Equal(suite.T(), moment, a2.Moment)
		assert.NoError(suite.T(), err, "Saved Activity (with moment) with no errors")
	})

	ts := httptest.NewServer(r)
	defer ts.Close()
	test.Request(suite.T(), ts, "GET", "/", nil)
}

func TestModelsSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
