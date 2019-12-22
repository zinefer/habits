package session_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/zinefer/habits/internal/habits/helpers/test"
	"github.com/zinefer/habits/internal/habits/middlewares/session"
)

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) TestDbContextMiddleware() {
	os := sessions.NewCookieStore([]byte("sessiontest"))
	r := chi.NewRouter()

	r.Use(session.SessionContextMiddleware(os))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		s := session.GetSessionFromContext(r.Context())
		assert.Equal(suite.T(), os, s)
	})

	ts := httptest.NewServer(r)
	defer ts.Close()

	test.Request(suite.T(), ts, "GET", "/", nil)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
